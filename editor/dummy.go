package editor

import (
	"fmt"
)

type DummyWidget struct {
}

func (v *DummyWidget) Sync() {

}

func (v *DummyWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *DummyWidget) Build() {
}

func (v *DummyWidget) Interface() interface{} {
	return nil
}

func Dummy() *DummyWidget {
	return &DummyWidget{}
}
