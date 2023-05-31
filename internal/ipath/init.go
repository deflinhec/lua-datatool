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

package ipath

import (
	"fmt"
	"math"
	"strings"
)

type Entry []string

func (v Entry) Filter(fn func(string) bool) Entry {
	entry := make(Entry, 0)
	for _, e := range v {
		if fn(e) {
			continue
		}
		entry = append(entry, e)
	}
	return entry
}

func (v Entry) Depth(d int) Entry {
	entry := make(Entry, 0)
	for _, e := range v {
		if strings.Count(e, ".") != d {
			continue
		}

	}
	return entry
}

func Walk(value interface{}) Entry {
	return walk(value, []string{})
}

func walk(value interface{}, paths []string) Entry {
	var parent string
	if len(paths) > 0 {
		parent = paths[len(paths)-1]
	}
	switch value := value.(type) {
	case map[string]interface{}:
		for k, v := range value {
			path := parent
			if len(path) > 0 {
				path += "."
			}
			path += k
			paths = walk(v, append(paths, path))
		}
	case []interface{}:
		for i, v := range value {
			path := parent
			path += fmt.Sprintf("[%v]", i)
			paths = walk(v, append(paths, path))
		}
	case string:
		path := parent
		if len(path) > 0 {
			path += "="
		}
		path += fmt.Sprint(`"` + value + `"`)
		paths = append(paths, path)
	default:
		path := parent
		if len(path) > 0 {
			path += "="
		}
		path += fmt.Sprint(value)
		paths = append(paths, path)
	}
	return paths
}

func Parent(value string) string {
	if len(value) == 0 {
		return ""
	}
	value = strings.Trim(value, ".")
	route := strings.Split(value, "=")
	if len(route) > 1 {
		route = strings.Split(route[len(route)-1], ".")
		return route[len(route)-1]
	} else if len(route) == 1 {
		return ""
	}
	route = strings.Split(route[len(route)-1], ".")
	return route[len(route)-1]
}

func Root(value string) string {
	value = strings.Trim(value, ".")
	route := strings.Split(value, ".")
	if len(route) > 0 {
		return route[0]
	}
	route = strings.Split(value, "=")
	if len(route) > 0 {
		return route[0]
	}
	return ""
}

func Spilt(value string) []string {
	value = strings.Trim(value, ".")
	route := strings.Split(value, ".")
	if len(route) > 0 {
		pair := strings.Split(route[len(route)-1], "=")
		if len(pair) > 1 {
			route[len(route)-1] = pair[0]
			route = append(route, pair[1])
		}
	}
	return route
}

func Depth(value string, d int) string {
	value = strings.Trim(value, ".")
	splits := Spilt(value)
	if len(splits) > d {
		return splits[d]
	}
	return ""
}

func Value(value string) string {
	value = strings.Trim(value, ".")
	splits := strings.Split(value, ".")
	if len(splits) == 0 {
		return ""
	}
	pair := strings.Split(splits[len(splits)-1], "=")
	if len(pair) > 1 {
		return pair[len(pair)-1]
	}
	return ""
}

func maxInt(value, max int) int {
	return int(math.Max(float64(value), float64(max)))
}

func Dir(value string, d int) string {
	value = strings.Trim(value, ".")
	splits := Spilt(value)
	if len(splits) == 0 {
		return ""
	}
	route := splits[0]
	d = maxInt(d, len(splits)-1)
	for i := 1; i < d; i++ {
		value += "." + splits[i]
	}
	return route
}

func Trim(value string, cutset string) string {
	value = strings.Trim(value, ".")
	value = strings.Trim(value, cutset)
	return strings.Trim(value, ".")
}
