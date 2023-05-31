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

package formatter

import (
	"fmt"
	"hash/fnv"
	"math"
	"sort"
	"strconv"
	"strings"

	zerowidth "github.com/trubitsyn/go-zero-width"
)

func Format(name string, value interface{}, opts ...Option) string {
	option := options{
		style: style{
			maxLineLength:   60,
			indentWidhtStep: 4,
			keySort:         KeySortAlphabetical,
			keySortPriority: []string{},
		},
	}
	for _, opt := range opts {
		opt(&option)
	}
	width, trait, style := option.rootIndentWidth, option.rootTrait, &option.style
	return fmt.Sprintf("%v=\n%v", name, format(value, width, trait, style))
}

func format(value interface{}, indent int, trait string, style *style) string {
	trait = strings.TrimPrefix(trait, ".")
	var ctx string
	switch value := value.(type) {
	case nil:
		ctx += "nil"
	case int:
		ctx += fmt.Sprint(value)
	case int8:
		ctx += fmt.Sprint(value)
	case int16:
		ctx += fmt.Sprint(value)
	case int32:
		ctx += fmt.Sprint(value)
	case int64:
		ctx += fmt.Sprint(value)
	case uint:
		ctx += fmt.Sprint(value)
	case uint8:
		ctx += fmt.Sprint(value)
	case uint16:
		ctx += fmt.Sprint(value)
	case uint32:
		ctx += fmt.Sprint(value)
	case uint64:
		ctx += fmt.Sprint(value)
	case float32:
		ctx += fmt.Sprint(value)
	case float64:
		ctx += fmt.Sprint(value)
	case bool:
		ctx += fmt.Sprint(value)
	case string:
		ctx += strconv.QuoteToGraphic(zerowidth.RemoveZeroWidthCharacters(value))
	case map[string]interface{}:
		keys := make(sort.StringSlice, 0, len(value))
		for k := range value {
			keys = append(keys, k)
		}
		switch style.keySort {
		case KeySortHashFnv64:
			sort.Slice(keys, func(i, j int) bool {
				return hashIndex(keys[i]) < hashIndex(keys[j])
			})
		case KeySortHashFnv64Reversed:
			sort.Slice(keys, func(i, j int) bool {
				return hashIndex(keys[i]) > hashIndex(keys[j])
			})
		case KeySortAlphabeticalReversed:
			sort.Sort(sort.Reverse(keys))
		case KeySortAlphabetical:
			sort.Strings(keys)
		}
		if len(style.keySortPriority) > 0 {
			sort.Slice(keys, func(i, j int) bool {
				l, r := math.MaxInt32, math.MaxInt32
				for k, v := range style.keySortPriority {
					if keys[i] == v {
						l = k
					}
					if keys[j] == v {
						r = k
					}
				}
				return l < r
			})
		}
		ctx += "{" + fmt.Sprintf(" -- %v", trait) + "\n"
		ctx += strings.Repeat(" ", indent+style.indentWidhtStep)
		for _, k := range keys {
			var subtrait, kv string
			switch k {
			case "true", "false":
				subtrait = fmt.Sprintf(`%v["%v"]`, trait, k)
				kv += fmt.Sprintf(`["%v"]=`, k)
			default:
				if strings.HasPrefix(k, string(KeyEscapeRune)) {
					subtrait = fmt.Sprintf("%v.%v", trait, k[1:])
					kv += fmt.Sprintf(`[%v]=`, k[1:])
				} else {
					subtrait = fmt.Sprintf("%v.%v", trait, k)
					kv += fmt.Sprintf(`%v=`, k)
				}
			}
			kv += format(value[k], indent+style.indentWidhtStep, subtrait, style) + ","
			kv += "\n" + strings.Repeat(" ", indent+style.indentWidhtStep)
			ctx += kv
		}
		if len(keys) > 0 {
			ctx = ctx[:len(ctx)-style.indentWidhtStep]
		}
		ctx += "}"
	case []interface{}:
		ctx += "{" + fmt.Sprintf(" -- total: %v", len(value))
		var line, chunk string
		for i, v := range value {
			key := i + 1
			subtrait := fmt.Sprintf("%v[%v]", trait, key)
			switch {
			case len(value) >= 25550:
				line += fmt.Sprintf("[%v]=%v,", key, format(v, indent+style.indentWidhtStep, subtrait, style))
			default:
				line += fmt.Sprintf("%v,", format(v, indent+style.indentWidhtStep, subtrait, style))
			}
			if len(line) > style.maxLineLength || key == len(value) {
				chunk += strings.Repeat(" ", indent+style.indentWidhtStep) + line + "\n"
				line = ""
			}
		}
		ctx += "\n" + align(chunk, indent, style) + strings.Repeat(" ", indent) + "}"
	default:
		panic(fmt.Sprintf("unsupported type: %T", value))
	}
	return ctx
}

// Align sequence chunk.
func align(chunk string, indent int, style *style) string {
	if strings.Contains(chunk, "{") {
		return chunk
	}
	var length int
	sections := make([]string, 0)
	chunk = strings.ReplaceAll(chunk, "\n", "")
	for _, section := range strings.Split(chunk, ",") {
		section = strings.TrimSpace(section)
		if len(section) == 0 {
			continue
		}
		sections = append(sections, section)
		length = int(math.Max(float64(length), float64(len(section))))
	}
	var line, format string
	format = fmt.Sprintf("%%%vs", length)
	line, chunk = strings.Repeat(" ", indent+style.indentWidhtStep), ""
	for i, section := range sections {
		line += fmt.Sprintf(format, section) + ","
		if len(line) > style.maxLineLength || i == len(sections)-1 {
			chunk += line + "\n"
			line = strings.Repeat(" ", indent+style.indentWidhtStep)
		}
	}
	if len(strings.Trim(line, " ")) > 0 {
		chunk += line + "\n"
	}
	return chunk
}

func hashIndex(index string) int64 {
	index = strings.TrimPrefix(index, string(KeyEscapeRune))
	i64, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		l := fnv.New64a()
		l.Write([]byte(index))
		i64 = int64(l.Sum64())
	}
	return i64
}
