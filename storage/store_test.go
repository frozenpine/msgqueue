package storage_test

import (
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/frozenpine/msgqueue/flow"
	"github.com/frozenpine/msgqueue/storage"
)

type Int struct {
	int
}

func (v Int) Len() int {
	return 4
}

func (v Int) Serialize() []byte {
	result := make([]byte, 4)

	binary.LittleEndian.PutUint32(result, uint32(v.int))

	return result
}

func (v *Int) Deserialize(data []byte) error {
	result := binary.LittleEndian.Uint32(data)
	v.int = int(result)
	return nil
}

type Varaint struct {
	name string
	data Int
}

func (v Varaint) Len() int {
	return -1
}

func (v Varaint) Serialize() []byte {
	result := []byte(v.name)

	result = append(result, v.data.Serialize()...)

	return result
}

func (v *Varaint) Deserialize(data []byte) error {
	len := len(data)
	v.data.Deserialize(data[len-4:])
	v.name = string(data[0 : len-4])
	return nil
}

func TestFileStore(t *testing.T) {
	flowFile := "flow.dat"

	store := storage.NewFileStore(flowFile)

	if err := store.Open(storage.WROnly); err != nil {
		t.Fatal("store open failed:", err)
	}

	t1 := flow.RegisterType(func() flow.PersistentData {
		return &Int{}
	})
	t2 := flow.RegisterType(func() flow.PersistentData {
		return &Varaint{}
	})

	v1 := Int{100}
	v2 := Varaint{
		name: "testtest",
		data: Int{200},
	}

	item1 := flow.FlowItem{
		Epoch:    0,
		Sequence: 1,
		TID:      t1,
		Data:     &v1,
	}

	item2 := flow.FlowItem{
		Epoch:    0,
		Sequence: 2,
		TID:      t2,
		Data:     &v2,
	}

	if err := store.Write(&item1); err != nil {
		t.Fatal(err)
	}

	if err := store.Write(&item2); err != nil {
		t.Fatal(err)
	}

	if err := store.Close(); err != nil {
		t.Fatal("store close failed:", err)
	}

	store = storage.NewFileStore(flowFile)
	if err := store.Open(storage.RDOnly); err != nil {
		t.Fatal("store open failed:", err)
	}

	if rd1, err := store.Read(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(rd1, *rd1.Data.(*Int))
	}

	if rd2, err := store.Read(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(rd2, *rd2.Data.(*Varaint))
	}

	if _, err := store.Read(); errors.Is(err, io.EOF) {
		t.Log("end of flow file")
	} else {
		t.Fatal(err)
	}

	store.Close()
}
