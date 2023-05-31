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
	"hash/fnv"
	"log"

	g "github.com/AllenDang/giu"
	"github.com/deflinhec/lua-datatool/editor/field"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/deflinhec/lua-datatool/internal/translations"
)

var translator *message.Printer

func init() {
	lang := language.MustParse("zh-TW")
	translator = message.NewPrinter(lang)
}

func Translator(p *message.Printer) {
	field.Translator(p)
	translator = p
}

type Widget interface {
	g.Widget
	Sync()
	String() string
	Interface() interface{}
}

func unmarshal(value interface{}) Widget {
	widget := field.Unmarshal(value)
	if widget != nil {
		return widget
	}
	switch v := value.(type) {
	case map[string]interface{}:
		widget = Table().Values(v).Editable()
	case []interface{}:
		widget = Slice().Values(v).Editable()
	default:
		log.Printf("not support type %T", v)
	}
	return widget
}

func Unmarshal(value interface{}) Widget {
	widget := unmarshal(value)
	if widget == nil {
		return nil
	}
	defer widget.Sync()
	return widget
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type AutoIDAllocator struct {
	ids map[string]string
}

func (v *AutoIDAllocator) Gen(id string) string {
	if id, ok := v.ids[id]; ok {
		return id
	}
	v.ids[id] = g.GenAutoID(id)
	return v.ids[id]
}

func AutoID() *AutoIDAllocator {
	return &AutoIDAllocator{ids: make(map[string]string)}
}

type StringSequence []string

func (s StringSequence) Cointains(e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s StringSequence) Index(e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func (s StringSequence) Erase(e int) StringSequence {
	if e >= len(s) {
		return s
	} else if e < 0 {
		return s
	}
	return append(s[:e], s[e+1:]...)
}

type WidgetSequence []Widget

func (s WidgetSequence) Erase(e int) WidgetSequence {
	if e >= len(s) {
		return s
	} else if e < 0 {
		return s
	}
	return append(s[:e], s[e+1:]...)
}

func (s WidgetSequence) Build() {
	for _, widget := range s {
		widget.Build()
	}
}

func (s WidgetSequence) Sync() {
	for _, widget := range s {
		widget.Sync()
	}
}
