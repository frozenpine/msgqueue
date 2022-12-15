package stream

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/frozenpine/msgqueue/core"
	"github.com/frozenpine/msgqueue/pipeline"
	"github.com/gofrs/uuid"
)

type DefaultWindow[
	IS, IV any,
	OS, OV any,
] struct {
	sequence []pipeline.Sequence[IS, IV]
}

func (win *DefaultWindow[IS, IV, OS, OV]) Indexs() []IS {
	index := make([]IS, len(win.sequence))

	for idx, v := range win.sequence {
		index[idx] = v.Index()
	}

	return index
}

func (win *DefaultWindow[IS, IV, OS, OV]) Values() []IV {
	values := make([]IV, len(win.sequence))

	for idx, v := range win.sequence {
		values[idx] = v.Value()
	}

	return values
}

func (win *DefaultWindow[IS, IV, OS, OV]) Series() []pipeline.Sequence[IS, IV] {
	return win.sequence
}

func (win *DefaultWindow[IS, IV, OS, OV]) Push(seq pipeline.Sequence[IS, IV]) error {
	if seq.IsWaterMark() {
		return ErrFutureTick
	}

	win.sequence = append(win.sequence, seq)

	return nil
}

func (win *DefaultWindow[IS, IV, OS, OV]) NextWindow() Window[IS, IV, OS, OV] {
	return &DefaultWindow[IS, IV, OS, OV]{}
}

type MemoStream[
	IS, IV any,
	OS, OV any,
	KEY comparable,
] struct {
	name     string
	id       uuid.UUID
	runCtx   context.Context
	cancelFn context.CancelFunc

	initOnce, releaseOnce sync.Once

	pipeline pipeline.Pipeline[IS, IV, OS, OV]

	windowCache []Window[IS, IV, OS, OV]
	currWindow  Window[IS, IV, OS, OV]

	aggregator Aggregator[IS, IV, OS, OV]
}

func NewMemoStream[
	IS, IV any,
	OS, OV any,
	KEY comparable,
](
	ctx context.Context, name string,
	initWin Window[IS, IV, OS, OV],
	agg Aggregator[IS, IV, OS, OV],
) (*MemoStream[IS, IV, OS, OV, KEY], error) {
	if agg == nil {
		return nil, errors.Wrap(ErrInvalidAggregator, "aggregator missing")
	}
	stream := MemoStream[IS, IV, OS, OV, KEY]{}

	stream.Init(ctx, name, func() {
		if initWin == nil {
			initWin = &DefaultWindow[IS, IV, OS, OV]{}
		}

		stream.pipeline = pipeline.NewMemoPipeLine(ctx, name, stream.convert)
		stream.currWindow = initWin
	})

	return &stream, nil
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) convert(inData pipeline.Sequence[IS, IV], outChan core.Producer[pipeline.Sequence[OS, OV]]) error {
	err := strm.currWindow.Push(inData)

	switch errors.Unwrap(err) {
	case ErrFutureTick:
		result, err := strm.aggregator(strm.currWindow)

		if err != nil {
			return err
		}

		if err = outChan.Publish(result, -1); err != nil {
			log.Printf("Stream out failed: +%v", err)
		}

		strm.windowCache = append(strm.windowCache, strm.currWindow)
		strm.currWindow = strm.currWindow.NextWindow()
	default:
		return err
	}

	return nil
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) ID() uuid.UUID {
	return strm.id
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Name() string {
	return strm.name
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Init(ctx context.Context, name string, extraInit func()) {
	strm.initOnce.Do(func() {
		if ctx == nil {
			ctx = context.Background()
		}

		if name == "" {
			name = "MemoStream"
		}

		strm.name = core.GenName(name)
		strm.id = core.GenID(strm.name)

		if extraInit != nil {
			extraInit()
		}
	})
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Join() {
	<-strm.runCtx.Done()

	strm.pipeline.Join()

	// TODO: extra join
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Release() {
	strm.releaseOnce.Do(func() {
		strm.cancelFn()

		strm.pipeline.Release()
		// TODO: extra release
	})
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Publish(v pipeline.Sequence[IS, IV], timeout time.Duration) error {
	return strm.pipeline.Publish(v, timeout)
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) Subscribe(name string, resume core.ResumeType) (uuid.UUID, <-chan pipeline.Sequence[OS, OV]) {
	return strm.pipeline.Subscribe(name, resume)
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) UnSubscribe(subID uuid.UUID) error {
	return strm.pipeline.UnSubscribe(subID)
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) PipelineUpstream(src core.Consumer[pipeline.Sequence[IS, IV]]) error {
	return strm.pipeline.PipelineUpStream(src)
}

func (strm *MemoStream[IS, IV, OS, OV, KEY]) PipelineDownStream(dst core.Upstream[pipeline.Sequence[OS, OV]]) error {
	return strm.pipeline.PipelineDownStream(dst)
}
