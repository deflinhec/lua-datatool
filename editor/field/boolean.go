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
	"fmt"

	g "github.com/AllenDang/giu"
)

type BooleanWidget struct {
	g.Widget
	Value       bool
	placeholder bool
}

func (v *BooleanWidget) Sync() {
	v.placeholder = v.Value
}

func (v *BooleanWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *BooleanWidget) Interface() interface{} {
	return v.Value
}

func (v *BooleanWidget) Build() {
	g.Checkbox("", &v.placeholder).
		OnChange(func() {
			v.Value = v.placeholder
		}).
		Build()
}

func Bool(value bool) *BooleanWidget {
	return &BooleanWidget{Value: value}
}
