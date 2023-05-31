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
	"strconv"

	g "github.com/AllenDang/giu"
)

type StringWidget struct {
	g.Widget
	placeholder string
	Value       string
}

func (v *StringWidget) Sync() {
	value := strconv.QuoteToGraphic(v.Value)
	v.placeholder = value[1 : len(value)-1]
}

func (v *StringWidget) String() string {
	return v.Value
}

func (v *StringWidget) Interface() interface{} {
	return v.Value
}

func (v *StringWidget) Build() {
	g.InputText(&v.placeholder).
		Size(192).
		OnChange(func() {
			value := `"` + v.placeholder + `"`
			s, err := strconv.Unquote(value)
			if err != nil {
				defer v.Sync()
				g.Msgbox(translator.Sprintf("錯誤"), err.Error()).
					Buttons(g.MsgboxButtonsOk)
				return
			}
			v.Value = s
		}).
		Build()
}

func String(value string) *StringWidget {
	return &StringWidget{Value: value}
}
