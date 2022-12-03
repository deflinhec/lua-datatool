package editor

import (
	"image"

	g "github.com/AllenDang/giu"
)

type IndexEditorWidget struct {
	g.Widget
	selected    int32
	placeholder string
	indexes     StringSequence
	types       []string
	alloc       *AutoIDAllocator
	onAdd       func(string, Widget)
	onRemove    func(int32)
}

func (v *IndexEditorWidget) clear() {
	v.selected, v.placeholder = -1, ""
}

func (v *IndexEditorWidget) OnAdd(fn func(string, Widget)) *IndexEditorWidget {
	v.onAdd = fn
	return v
}

func (v *IndexEditorWidget) OnRemove(fn func(int32)) *IndexEditorWidget {
	v.onRemove = fn
	return v
}

func (v *IndexEditorWidget) Indexes(indexes ...string) *IndexEditorWidget {
	v.indexes = indexes
	return v
}

func (v *IndexEditorWidget) Build() {
	g.Row(
		g.Condition(v.onAdd != nil, g.Layout{
			g.Custom(func() {
				id := v.alloc.Gen("新增欄位")
				g.SetNextWindowSize(300, 210)
				g.PopupModal(id).
					Layout(
						g.Row(
							g.Label(translator.Sprintf("欄位名稱：")),
							g.InputText(&v.placeholder).
								Size(192).
								Flags(g.InputTextFlagsCharsNoBlank|
									g.InputTextFlagsEnterReturnsTrue|
									g.InputTextFlagsAutoSelectAll),
						),
						g.Row(
							g.Label(translator.Sprintf("數值類型：")),
							g.ListBox(v.alloc.Gen(id), v.types).
								Size(192, 90).
								OnChange(func(selected int) {
									v.selected = int32(selected)
								}),
						),
						g.Row(
							g.Button(translator.Sprintf("新增")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
									if len(v.placeholder) == 0 {
										g.Msgbox(translator.Sprintf("警告"),
											translator.Sprintf("欄位名稱不能為空")).
											Buttons(g.MsgboxButtonsOk)
										return
									}
									if v.indexes.Cointains(v.placeholder) {
										g.Msgbox(translator.Sprintf("警告"),
											translator.Sprintf("欄位名稱重複")).
											Buttons(g.MsgboxButtonsOk)
										return
									}
									var value Widget
									switch v.types[v.selected] {
									case "string":
										value = unmarshal("")
									case "number":
										value = unmarshal(float64(0))
									case "table":
										value = unmarshal(map[string]interface{}{})
									case "slice":
										value = unmarshal([]interface{}{})
									}
									v.onAdd(v.placeholder, value)
								}),
							g.Button(translator.Sprintf("取消")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
								}),
						),
					).
					Build()
				g.Button(translator.Sprintf("新增欄位")).
					OnClick(func() {
						v.selected = 0
						g.OpenPopup(id)
					}).
					Build()
			}),
		}, g.Layout{}),
		g.Condition(v.onRemove != nil, g.Layout{
			g.Custom(func() {
				id := v.alloc.Gen("移除欄位")
				g.SetNextWindowSize(300, 210)
				g.PopupModal(id).
					Layout(
						g.ListBox(g.GenAutoID(id), v.indexes).
							Size(300, 80).
							OnChange(func(selected int) {
								v.selected = int32(selected)
							}),
						g.Row(
							g.Button(translator.Sprintf("移除")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
									v.onRemove(v.selected)
								}),
							g.Button(translator.Sprintf("取消")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
								}),
						),
					).
					Build()
				g.Button(translator.Sprintf("移除欄位")).
					Disabled(len(v.indexes) == 0).
					OnClick(func() {
						v.selected = 0
						g.OpenPopup(id)
					}).
					Build()
			}),
		}, g.Layout{}),
	).Build()
}

func IndexEditor() *IndexEditorWidget {
	return &IndexEditorWidget{
		indexes: make(StringSequence, 0),
		types:   []string{"string", "number", "table", "slice"},
		alloc:   AutoID(),
	}
}

type SliceEditorWidget struct {
	g.Widget
	index    int32
	selected int32
	indexes  StringSequence
	types    []string
	alloc    *AutoIDAllocator
	onInsert func(int32, Widget)
	onRemove func(int32)
}

func (v *SliceEditorWidget) clear() {
	v.selected, v.index = -1, 0
}

func (v *SliceEditorWidget) OnInsert(fn func(int32, Widget)) *SliceEditorWidget {
	v.onInsert = fn
	return v
}

func (v *SliceEditorWidget) OnRemove(fn func(int32)) *SliceEditorWidget {
	v.onRemove = fn
	return v
}

func (v *SliceEditorWidget) Indexes(indexes ...string) *SliceEditorWidget {
	v.indexes = indexes
	return v
}

func (v *SliceEditorWidget) Build() {
	g.Row(
		g.Condition(v.onInsert != nil, g.Layout{
			g.Custom(func() {
				id := v.alloc.Gen("插入索引")
				g.SetNextWindowSize(300, 200)
				g.PopupModal(id).
					Layout(
						g.Condition(len(v.indexes) > 0, g.Layout{
							g.Row(
								g.Label(translator.Sprintf("索引序次：")),
								g.SliderInt(&v.index, 0, int32(len(v.indexes))),
							),
						}, g.Layout{}),
						g.Row(
							g.Label(translator.Sprintf("數值類型：")),
							g.ListBox(v.alloc.Gen(id), v.types).
								Size(192, 90).
								OnChange(func(selected int) {
									v.selected = int32(selected)
								}),
						),
						g.Row(
							g.Button(translator.Sprintf("插入")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
									var value Widget
									switch v.types[v.selected] {
									case "string":
										value = unmarshal("")
									case "number":
										value = unmarshal(float64(0))
									case "table":
										value = unmarshal(map[string]interface{}{})
									case "slice":
										value = unmarshal([]interface{}{})
									}
									v.onInsert(v.index, value)
								}),
							g.Button(translator.Sprintf("取消")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
								}),
						),
					).
					Build()
				g.Button(translator.Sprintf("插入索引")).
					OnClick(func() {
						v.selected = 0
						g.OpenPopup(id)
					}).
					Build()
			}),
		}, g.Layout{}),
		g.Condition(v.onRemove != nil, g.Layout{
			g.Custom(func() {
				id := v.alloc.Gen("移除索引")
				g.SetNextWindowSize(300, 150)
				g.PopupModal(id).
					Layout(
						g.ListBox(g.GenAutoID(id), v.indexes).
							Size(300, 80).
							OnChange(func(selected int) {
								v.selected = int32(selected)
							}),
						g.Row(
							g.Button(translator.Sprintf("移除")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
									v.onRemove(v.selected)
								}),
							g.Button(translator.Sprintf("取消")).
								OnClick(func() {
									defer v.clear()
									defer g.CloseCurrentPopup()
								}),
						),
					).
					Build()
				g.Button(translator.Sprintf("移除索引")).
					Disabled(len(v.indexes) == 0).
					OnClick(func() {
						v.selected = 0
						g.OpenPopup(id)
					}).
					Build()
			}),
		}, g.Layout{}),
	).Build()
}

func SliceEditor() *SliceEditorWidget {
	return &SliceEditorWidget{
		index:   0,
		indexes: make(StringSequence, 0),
		types:   []string{"string", "number", "table", "slice"},
		alloc:   AutoID(),
	}
}

type IndexSelectorWidget struct {
	g.Widget
	selected    int32
	placeholder string
	searching   string
	hash        map[string]bool
	indexes     StringSequence
	onSelected  func(string)
}

func (v *IndexSelectorWidget) OnSelected(fn func(string)) *IndexSelectorWidget {
	v.onSelected = fn
	return v
}

func (v *IndexSelectorWidget) Indexes(indexes ...string) *IndexSelectorWidget {
	v.indexes = indexes
	v.hash = make(map[string]bool, 0)
	for _, index := range v.indexes {
		v.hash[index] = true
	}
	return v
}

func (v *IndexSelectorWidget) Build() {
	g.Condition(v.onSelected != nil, g.Layout{
		g.Row(
			g.Label(translator.Sprintf("索引")),
			g.Combo("", v.placeholder, v.indexes, &v.selected).
				Size(200).
				OnChange(func() {
					v.placeholder = v.indexes[v.selected]
					v.onSelected(v.placeholder)
				}),
			g.InputText(&v.searching).
				Size(200).
				Hint(translator.Sprintf("檢索字元")).
				AutoComplete(v.indexes).
				Flags(g.InputTextFlagsAutoSelectAll|
					g.InputTextFlagsCharsNoBlank).
				OnChange(func() {
					if b, ok := v.hash[v.searching]; ok && b {
						defer func() { v.searching = "" }()
						v.placeholder = v.searching
						v.onSelected(v.placeholder)
					}
				}),
			g.Label(translator.Sprintf("總計：%v", len(v.indexes))),
		),
	}, g.Layout{}).Build()
}

func IndexSelector() *IndexSelectorWidget {
	return &IndexSelectorWidget{
		indexes: make(StringSequence, 0),
		hash:    make(map[string]bool, 0),
	}
}

type IndexButtonWidget struct {
	index   int32
	title   string
	onClick func(int32)
}

func (v *IndexButtonWidget) Build() {
	g.Button(v.title).
		OnClick(func() {
			if v.onClick != nil {
				v.onClick(v.index)
			}
		}).
		Build()
}

func (v *IndexButtonWidget) Title(title string) *IndexButtonWidget {
	v.title = title
	return v
}

func (v *IndexButtonWidget) Index(index int32) *IndexButtonWidget {
	v.index = index
	return v
}

func (v *IndexButtonWidget) OnClick(fn func(int32)) *IndexButtonWidget {
	v.onClick = fn
	return v
}

func IndexButton(title string) *IndexButtonWidget {
	return &IndexButtonWidget{
		title: title,
		index: -1,
	}
}

type IndexImageButtonWidget struct {
	index   int32
	image   *image.RGBA
	onClick func(int32)
}

func (v *IndexImageButtonWidget) Build() {
	g.ImageButtonWithRgba(v.image).
		Size(16, 16).
		OnClick(func() {
			if v.onClick != nil {
				v.onClick(v.index)
			}
		}).
		Build()
}

func (v *IndexImageButtonWidget) Index(index int32) *IndexImageButtonWidget {
	v.index = index
	return v
}

func (v *IndexImageButtonWidget) OnClick(fn func(int32)) *IndexImageButtonWidget {
	v.onClick = fn
	return v
}

func IndexImageButton(name string) *IndexImageButtonWidget {
	return &IndexImageButtonWidget{
		image: Image(name),
	}
}

type IndexFilterWidget struct {
	visible map[string]bool
	indexes StringSequence
}

func (v *IndexFilterWidget) Indexes(indexes ...string) *IndexFilterWidget {
	v.indexes = indexes
	return v
}

func IndexFilter() *IndexFilterWidget {
	return &IndexFilterWidget{
		visible: make(map[string]bool),
	}
}
