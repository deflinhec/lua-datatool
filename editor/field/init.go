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
	g "github.com/AllenDang/giu"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var translator *message.Printer

func init() {
	lang := language.MustParse("zh-TW")
	translator = message.NewPrinter(lang)
}

func Translator(p *message.Printer) {
	translator = p
}

type Widget interface {
	g.Widget
	Sync()
	String() string
	Interface() interface{}
}

func Unmarshal(value interface{}) Widget {
	var widget Widget
	switch v := value.(type) {
	case int:
		widget = Int(v)
	case int16:
		widget = Int16(v)
	case int32:
		widget = Int32(v)
	case int64:
		widget = Int64(v)
	case uint:
		widget = Uint(v)
	case uint16:
		widget = Uint16(v)
	case uint32:
		widget = Uint32(v)
	case uint64:
		widget = Uint64(v)
	case string:
		widget = String(v)
	case float32:
		widget = Float32(v)
	case float64:
		widget = Float64(v)
	case bool:
		widget = Bool(v)
	}
	return widget
}
