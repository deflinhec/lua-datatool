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
	"sort"

	g "github.com/AllenDang/giu"
)

type ScopeWidget struct {
	keys     StringSequence
	visible  WidgetSequence
	datas    map[string]Widget
	selector *IndexSelectorWidget
	editor   *IndexEditorWidget
	editable bool
}

func (v *ScopeWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *ScopeWidget) Sync() {
	sort.Strings(v.keys)
	v.selector.Indexes(v.keys...)
	v.editor.Indexes(v.keys...)
	v.visible.Sync()
}

func (v *ScopeWidget) Interface() interface{} {
	table := make(map[string]interface{})
	for _, key := range v.keys {
		table[key] = v.datas[key].Interface()
	}
	return table
}

func (v *ScopeWidget) Build() {
	g.Row(
		g.Condition(v.editable, g.Layout{
			v.editor.
				OnAdd(func(key string, value Widget) {
					v.keys = append(v.keys, key)
					v.datas[key] = value
					v.Sync()
				}).
				OnRemove(func(i int32) {
					delete(v.datas, v.keys[i])
					v.keys = v.keys.Erase(int(i))
					v.Sync()
				}),
		}, g.Layout{}),
		v.selector.
			OnSelected(func(selected string) {
				v.visible = make(WidgetSequence, 0, 1)
				if data, ok := v.datas[selected]; ok {
					v.visible = append(v.visible, data)
				}
				v.Sync()
			})).
		Build()
	g.Separator().
		Build()
	g.Child().
		Layout(v.visible).
		Build()
}

func (v *ScopeWidget) Len() int {
	return len(v.datas)
}

func (v *ScopeWidget) Editable() *ScopeWidget {
	v.editable = true
	return v
}

func (v *ScopeWidget) Values(table map[string]interface{}) *ScopeWidget {
	defer v.Sync()
	v.keys = make([]string, 0)
	v.visible = make(WidgetSequence, 0)
	v.datas = make(map[string]Widget, 0)
	for key, value := range table {
		v.keys = append(v.keys, fmt.Sprint(key))
		v.datas[key] = unmarshal(value)
	}
	return v
}

func Scope() *ScopeWidget {
	v := &ScopeWidget{
		keys:     make(StringSequence, 0),
		visible:  make(WidgetSequence, 0),
		datas:    make(map[string]Widget, 0),
		selector: IndexSelector(),
		editor:   IndexEditor(),
	}
	return v
}
