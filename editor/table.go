package editor

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	g "github.com/AllenDang/giu"
)

type TableWidget struct {
	indexes  []uint32
	keys     StringSequence
	values   map[string]Widget
	table    *g.TreeTableWidget
	editor   *IndexEditorWidget
	editable bool
}

func (v *TableWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *TableWidget) Sync() {
	sort.Sort(v)
	rows := make([]*g.TreeTableRowWidget, 0)
	for i, key := range v.keys {
		defer v.values[key].Sync()
		widget := g.Widget(v.values[key])
		if v.editable {
			widget = g.Row(widget,
				IndexImageButton("assets/bin.png").
					Index(int32(i)).
					OnClick(func(i int32) {
						key = v.keys[i]
						delete(v.values, key)
						v.keys = v.keys.Erase(int(i))
						v.Sync()
					}),
			)
		}
		rows = append(rows, g.TreeTableRow(key, widget).
			Flags(g.TreeNodeFlagsSpanAvailWidth))
	}
	v.editor.Indexes(v.keys...)
	v.table.Rows(rows...)
}

func (v *TableWidget) Interface() interface{} {
	table := make(map[string]interface{})
	for _, key := range v.keys {
		table[key] = v.values[key].Interface()
	}
	return table
}

func (v *TableWidget) Build() {
	g.Row(
		g.Condition(v.editable, g.Layout{
			v.editor.
				OnAdd(func(key string, value Widget) {
					v.keys = append(v.keys, key)
					v.values[key] = value
					v.Sync()
				}).
				OnRemove(func(i int32) {
					key := v.keys[i]
					delete(v.values, key)
					v.keys = v.keys.Erase(int(i))
					v.Sync()
				}),
		}, g.Layout{})).
		Build()
	g.Separator().Build()
	v.table.Columns(
		g.TableColumn(translator.Sprintf("欄位")).
			Flags(g.TableColumnFlagsWidthFixed|
				g.TableColumnFlagsNoClip),
		g.TableColumn(translator.Sprintf("數值")).
			Flags(g.TableColumnFlagsWidthStretch|
				g.TableColumnFlagsNoClip),
	)
	v.table.Build()
}

func (v *TableWidget) Len() int {
	return len(v.keys)
}

func (v *TableWidget) Less(i, j int) bool {
	return v.indexes[i] < v.indexes[j]
}

func (v *TableWidget) Swap(i, j int) {
	key, idx := v.keys[i], v.indexes[i]
	v.keys[i], v.indexes[i] = v.keys[j], v.indexes[j]
	v.keys[j], v.indexes[i] = key, idx
}

func (v *TableWidget) Editable() *TableWidget {
	v.editable = true
	return v
}

func (v *TableWidget) Values(table map[string]interface{}) *TableWidget {
	v.indexes = make([]uint32, 0)
	v.keys = make(StringSequence, 0)
	v.values = make(map[string]Widget, 0)
	for key, value := range table {
		v.keys = append(v.keys, key)
		v.values[key] = unmarshal(value)
		idx := int(hash(key))
		if strings.HasPrefix(key, "#") {
			idx, _ = strconv.Atoi(key[1:])
		}
		v.indexes = append(v.indexes, uint32(idx))
	}
	return v
}

func Table() *TableWidget {
	return &TableWidget{
		indexes: make([]uint32, 0),
		keys:    make(StringSequence, 0),
		values:  make(map[string]Widget, 0),
		editor:  IndexEditor(),
		table:   g.TreeTable(),
	}
}
