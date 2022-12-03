package editor

import (
	"testing"
)

func TestSlice(t *testing.T) {
	if Slice().Values([]interface{}{
		"string", 2.2, 3, true,
	}).Editable() == nil {
		t.Fail()
	}
}

func TestTable(t *testing.T) {
	if Table().Values(map[string]interface{}{
		"string":  "depth.2",
		"number":  1,
		"float":   2.3,
		"integer": 2.3,
		"array2": []interface{}{
			"string", 2.2, 3, true,
		},
	}) == nil {
		t.Fail()
	}
}

func TestUnmarshal(t *testing.T) {
	if Unmarshal(map[string]interface{}{
		"#0": map[string]interface{}{
			"string":  "depth=1",
			"number":  1,
			"float":   2.3,
			"integer": 2.3,
			"array1": []interface{}{
				1, 2, 3, 4, 5, 6,
			},
			"table1": map[string]interface{}{
				"string":  "depth.2",
				"number":  1,
				"float":   2.3,
				"integer": 2.3,
				"array2": []interface{}{
					1.1, 2.2, 3.3, 4.4, 5.5, 6.6,
				},
				"table2": map[string]interface{}{
					"string":  "depth.3",
					"number":  1,
					"float":   2.3,
					"integer": 2.3,
					"array3": []interface{}{
						true, false, true,
					},
					"table3": map[string]interface{}{
						"string":  "depth=4",
						"number":  1,
						"float":   2.3,
						"integer": 2.3,
					},
				},
			},
		},
		"#2": map[string]interface{}{
			"string":  "depth=1",
			"number":  1,
			"float":   2.3,
			"integer": 2.3,
			"array1": []interface{}{
				1, 2, 3, 4, 5, 6,
			},
			"table1": map[string]interface{}{
				"string":  "depth.2",
				"number":  1,
				"float":   2.3,
				"integer": 2.3,
				"array2": []interface{}{
					1.1, 2.2, 3.3, 4.4, 5.5, 6.6,
				},
				"table2": map[string]interface{}{
					"string":  "depth.3",
					"number":  1,
					"float":   2.3,
					"integer": 2.3,
					"array3": []interface{}{
						true, false, true,
					},
					"table3": map[string]interface{}{
						"string":  "depth=4",
						"number":  1,
						"float":   2.3,
						"integer": 2.3,
					},
				},
			},
		},
	}) == nil {
		t.Fail()
	}
}
