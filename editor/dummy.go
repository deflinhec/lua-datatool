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
	"fmt"
)

type DummyWidget struct {
}

func (v *DummyWidget) Sync() {

}

func (v *DummyWidget) String() string {
	return fmt.Sprint(v.Interface())
}

func (v *DummyWidget) Build() {
}

func (v *DummyWidget) Interface() interface{} {
	return nil
}

func Dummy() *DummyWidget {
	return &DummyWidget{}
}
