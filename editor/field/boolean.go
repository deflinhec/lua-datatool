package field

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

type BooleanWidget struct {
	g.Widget
	Value       bool
	placeholder bool
}

func (v *BooleanWidget) Sync() {
	v.placeholder = v.Value
}

func (v *BooleanWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *BooleanWidget) Interface() interface{} {
	return v.Value
}

func (v *BooleanWidget) Build() {
	g.Checkbox("", &v.placeholder).
		OnChange(func() {
			v.Value = v.placeholder
		}).
		Build()
}

func Bool(value bool) *BooleanWidget {
	return &BooleanWidget{Value: value}
}
