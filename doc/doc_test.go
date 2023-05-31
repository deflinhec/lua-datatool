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

package doc_test

import (
	"testing"

	"github.com/deflinhec/lua-datatool/doc"
	"github.com/go-test/deep"
)

func TestFileIO(t *testing.T) {
	f, err := doc.OpenFile("tests/edge.lua")
	if err != nil {
		t.Error(err)
	}
	err = f.SaveTo("tests/edge.out.lua")
	if err != nil {
		t.Error(err)
	}
	lf, err := doc.OpenFile("tests/edge.lua")
	if err != nil {
		t.Error(err)
	}
	rf, err := doc.OpenFile("tests/edge.out.lua")
	if err != nil {
		t.Error(err)
	}
	if diff := deep.Equal(lf.Value, rf.Value); diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
