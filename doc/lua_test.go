package doc

import (
	"reflect"
	"testing"
)

func TestFileIO(t *testing.T) {
	f, err := Open("lua_test.lua")
	if err != nil {
		t.Error(err)
	}
	a, err := Read(f)
	if err != nil {
		t.Error(err)
	}
	f, err = Open("lua_test_out.lua", WithDocument(f))
	err = Write(f, a)
	if err != nil {
		t.Error(err)
	}
	f, err = Open("lua_test_out.lua")
	if err != nil {
		t.Error(err)
	}
	b, err := Read(f)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}
