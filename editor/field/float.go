package field

import (
	"fmt"
	"strconv"

	g "github.com/AllenDang/giu"
)

type Float32Widget struct {
	g.Widget
	placeholder string
	Value       float32
}

func (v *Float32Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Float32Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Float32Widget) Interface() interface{} {
	return v.Value
}

func (v *Float32Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			f32, err := strconv.ParseFloat(v.placeholder, 32)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = float32(f32)
		}).
		Build()
}

func Float32(value float32) *Float32Widget {
	return &Float32Widget{Value: value}
}

type Float64Widget struct {
	g.Widget
	placeholder string
	Value       float64
}

func (v *Float64Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Float64Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Float64Widget) Interface() interface{} {
	return v.Value
}

func (v *Float64Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			f64, err := strconv.ParseFloat(v.placeholder, 64)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = f64
		}).
		Build()
}

func Float64(value float64) *Float64Widget {
	return &Float64Widget{Value: value}
}
