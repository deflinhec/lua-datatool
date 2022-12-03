package ipath

import (
	"log"
	"strconv"
	"testing"
)

func TestExport(t *testing.T) {
	paths := Walk(map[string]interface{}{
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
	})
	for _, path := range paths {
		log.Println(strconv.Quote(path))
		// log.Println(Value(path))
	}
}
