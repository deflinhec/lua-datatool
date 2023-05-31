package doc_test

import (
	"testing"

	"datatool.deflinhec.dev/doc"
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
