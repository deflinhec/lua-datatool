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

package giuext

import (
	"log"

	g "github.com/AllenDang/giu"
)

type ComboTextWidget struct {
	g.Widget
	selected   int32
	list       []string
	label      string
	onSelected func(string)
}

func (w *ComboTextWidget) Build() {
	preview := ""
	if len(w.list) > 0 {
		preview = w.list[0]
	}
	g.Row(
		g.Label(w.label),
		g.Combo("", preview, w.list[1:], &w.selected).
			OnChange(func() {
				if w.onSelected == nil {
					return
				} else if w.selected < 0 {
					return
				}
				i := w.selected + 1
				value := w.list[i]
				w.list[i] = w.list[0]
				w.list[0] = value
				log.Print(value)
				w.onSelected(value)
			}),
	).Build()
}

func (w *ComboTextWidget) OnSelected(fn func(string)) *ComboTextWidget {
	w.onSelected = fn
	return w
}

func ComboText(label string, list []string) *ComboTextWidget {
	return &ComboTextWidget{
		selected: int32(len(list) - 1),
		label:    label,
		list:     list,
	}
}
