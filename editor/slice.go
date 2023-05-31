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

package editor

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

type SliceWidget struct {
	indexes  StringSequence
	values   WidgetSequence
	table    *g.TreeTableWidget
	editor   *SliceEditorWidget
	editable bool
}

func (v *SliceWidget) Sync() {
	v.indexes = make([]string, 0, len(v.values))
	for i := range v.values {
		v.indexes = append(v.indexes, fmt.Sprint(i))
	}
	rows := make([]*g.TreeTableRowWidget, 0)
	for i, value := range v.values {
		defer value.Sync()
		widget := g.Widget(v.values[i])
		if v.editable {
			widget = g.Row(widget,
				IndexImageButton("bin.png").
					Index(int32(i)).
					OnClick(func(i int32) {
						v.values = v.values.Erase(int(i))
						v.Sync()
					}),
			)
		}
		rows = append(rows, g.TreeTableRow(fmt.Sprint(i), widget).
			Flags(g.TreeNodeFlagsSpanAvailWidth))
	}
	v.table.Rows(rows...)
	v.editor.Indexes(v.indexes...)
}

func (v *SliceWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *SliceWidget) Interface() interface{} {
	slice := make([]interface{}, 0)
	for _, e := range v.values {
		slice = append(slice, e.Interface())
	}
	return slice
}

func (v *SliceWidget) Build() {
	g.Row(
		g.Condition(v.editable, g.Layout{
			v.editor.
				OnInsert(func(i int32, value Widget) {
					if i == int32(len(v.values)) {
						v.values = append(v.values, value)
					} else {
						v.values = append(v.values[:i+1], v.values[i:]...)
						v.values[i] = value
					}
					v.Sync()
				}).
				OnRemove(func(i int32) {
					v.values = v.values.Erase(int(i))
					v.Sync()
				}),
		}, g.Layout{})).
		Build()
	g.Separator().Build()
	v.table.Columns(
		g.TableColumn(translator.Sprintf("索引")).
			Flags(g.TableColumnFlagsWidthFixed|
				g.TableColumnFlagsNoClip),
		g.TableColumn(translator.Sprintf("數值")).
			Flags(g.TableColumnFlagsWidthStretch|
				g.TableColumnFlagsNoClip),
	)
	v.table.Build()
}

func (v *SliceWidget) Len() int {
	return len(v.values)
}

func (v *SliceWidget) Editable() *SliceWidget {
	v.editable = true
	return v
}

func (v *SliceWidget) Values(slice []interface{}) *SliceWidget {
	v.values = make(WidgetSequence, 0)
	for _, value := range slice {
		v.values = append(v.values, unmarshal(value))
	}
	return v
}

func Slice() *SliceWidget {
	return &SliceWidget{
		values:  make(WidgetSequence, 0),
		indexes: make(StringSequence, 0),
		editor:  SliceEditor(),
		table:   g.TreeTable(),
	}
}
