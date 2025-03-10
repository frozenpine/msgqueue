package chanio

import (
	"fmt"
	"log/slog"
	"reflect"
	"runtime"
	"sync"
)

type PersistentData interface {
	Serialize() []byte
	Deserialize([]byte) error
}

var (
	typeList  []sync.Pool
	typeCache map[reflect.Type]TID
)

type TID int

// RegisterType register PersistentData type.
// TID present unique identity for PersistentData
// and type register action must always in same order
func RegisterType(t PersistentData, newFn func() PersistentData) (tid TID) {
	if typeCache == nil {
		typeCache = make(map[reflect.Type]TID)
	}

	typ := reflect.Indirect(reflect.ValueOf(t)).Type()

	var exist bool

	if tid, exist = typeCache[typ]; !exist {
		typeList = append(typeList, sync.Pool{New: func() any { return newFn() }})
		tid = TID(len(typeList) - 1)
		typeCache[typ] = tid
		slog.Info(
			"type registered",
			slog.String("type", typ.String()),
			slog.Int("tid", int(tid)),
		)
	}

	return
}

// NewTypeValue create PersistentData type.
// TID is unique identity for type creation
// from persistent storage
func NewTypeValue(tid TID) (PersistentData, error) {
	if tid < 0 || int(tid) >= len(typeList) {
		return nil, fmt.Errorf("TID[%d] out of range", tid)
	}

	data := typeList[tid].Get().(PersistentData)

	// RAII for put back data to pool
	runtime.SetFinalizer(data, typeList[tid].Put)

	return data, nil
}

type BaseIO interface {
	Open(int) error
	Close() error
	Flush() error
}

type ChanIO interface {
	BaseIO

	Write(TID, PersistentData) error
	Read() (PersistentData, error)
}
