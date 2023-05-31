// Copyright 2023 Deflinhec
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
