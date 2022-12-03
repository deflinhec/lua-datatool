package field

import (
	"fmt"
	"strconv"

	g "github.com/AllenDang/giu"
)

type IntWidget struct {
	g.Widget
	placeholder string
	Value       int
}

func (v *IntWidget) Sync() {
	v.placeholder = v.String()
}

func (v *IntWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *IntWidget) Interface() interface{} {
	return v.Value
}

func (v *IntWidget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i32, err := strconv.ParseInt(v.placeholder, 0, 32)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = int(i32)
		}).
		Build()
}

func Int(value int) *IntWidget {
	return &IntWidget{Value: value}
}

type Int16Widget struct {
	g.Widget
	placeholder string
	Value       int16
}

func (v *Int16Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Int16Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Int16Widget) Interface() interface{} {
	return v.Value
}

func (v *Int16Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i16, err := strconv.ParseInt(v.placeholder, 0, 16)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = int16(i16)
		}).
		Build()
}

func Int16(value int16) *Int16Widget {
	return &Int16Widget{Value: value}
}

type Int32Widget struct {
	g.Widget
	placeholder string
	Value       int32
}

func (v *Int32Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Int32Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Int32Widget) Interface() interface{} {
	return v.Value
}

func (v *Int32Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i32, err := strconv.ParseInt(v.placeholder, 0, 32)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = int32(i32)
		}).
		Build()
}

func Int32(value int32) *Int32Widget {
	return &Int32Widget{Value: value}
}

type Int64Widget struct {
	g.Widget
	placeholder string
	Value       int64
}

func (v *Int64Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Int64Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Int64Widget) Interface() interface{} {
	return v.Value
}

func (v *Int64Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i64, err := strconv.ParseInt(v.placeholder, 0, 64)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = i64
		}).
		Build()
}

func Int64(value int64) *Int64Widget {
	return &Int64Widget{Value: value}
}

type UintWidget struct {
	g.Widget
	placeholder string
	Value       uint
}

func (v *UintWidget) Sync() {
	v.placeholder = v.String()
}

func (v *UintWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *UintWidget) Interface() interface{} {
	return v.Value
}

func (v *UintWidget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i32, err := strconv.ParseUint(v.placeholder, 0, 32)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = uint(i32)
		}).
		Build()
}

func Uint(value uint) *UintWidget {
	return &UintWidget{Value: value}
}

type Uint16Widget struct {
	g.Widget
	placeholder string
	Value       uint16
}

func (v *Uint16Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Uint16Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Uint16Widget) Interface() interface{} {
	return v.Value
}

func (v *Uint16Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i16, err := strconv.ParseUint(v.placeholder, 0, 16)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = uint16(i16)
		}).
		Build()
}

func Uint16(value uint16) *Uint16Widget {
	return &Uint16Widget{Value: value}
}

type Uint32Widget struct {
	g.Widget
	placeholder string
	Value       uint32
}

func (v *Uint32Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Uint32Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Uint32Widget) Interface() interface{} {
	return v.Value
}

func (v *Uint32Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i32, err := strconv.ParseUint(v.placeholder, 0, 32)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = uint32(i32)
		}).
		Build()
}

func Uint32(value uint32) *Uint32Widget {
	return &Uint32Widget{Value: value}
}

type Uint64Widget struct {
	g.Widget
	placeholder string
	Value       uint64
}

func (v *Uint64Widget) Sync() {
	v.placeholder = v.String()
}

func (v *Uint64Widget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *Uint64Widget) Interface() interface{} {
	return v.Value
}

func (v *Uint64Widget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		Flags(g.InputTextFlagsCharsDecimal).
		OnChange(func() {
			i64, err := strconv.ParseUint(v.placeholder, 0, 64)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = i64
		}).
		Build()
}

func Uint64(value uint64) *Uint64Widget {
	return &Uint64Widget{Value: value}
}
