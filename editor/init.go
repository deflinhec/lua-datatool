package editor

import (
	"hash/fnv"
	"log"

	"datatool.deflinhec.dev/editor/field"
	g "github.com/AllenDang/giu"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "datatool.deflinhec.dev/internal/translations"
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
