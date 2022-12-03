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
