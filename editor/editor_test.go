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
