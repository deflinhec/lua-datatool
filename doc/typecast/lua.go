package typecast

import (
	"fmt"
	"time"

	zerowidth "github.com/trubitsyn/go-zero-width"
	lua "github.com/yuin/gopher-lua"
)

func RuntimeLuaConvertLuaValue(lv lua.LValue) interface{} {
	// Taken from: https://github.com/yuin/gluamapper/blob/master/gluamapper.go#L79
	switch v := lv.(type) {
	case *lua.LNilType:
		return nil
	case lua.LBool:
		return bool(v)
	case lua.LString:
		return zerowidth.RemoveZeroWidthCharacters(string(v))
	case lua.LNumber:
		vf := float64(v)
		vi := int64(v)
		if vf == float64(vi) {
			// If it's a whole number use an actual integer type.
			return vi
		}
		return vf
	case *lua.LTable:
		maxn := v.MaxN()
		if maxn == 0 {
			// Table.
			ret := make(map[string]interface{})
			v.ForEach(func(key, value lua.LValue) {
				keyStr := fmt.Sprint(RuntimeLuaConvertLuaValue(key))
				ret[keyStr] = RuntimeLuaConvertLuaValue(value)
			})
			return ret
		}
		// Array.
		ret := make([]interface{}, 0, maxn)
		for i := 1; i <= maxn; i++ {
			ret = append(ret, RuntimeLuaConvertLuaValue(v.RawGetInt(i)))
		}
		return ret
	case *lua.LFunction:
		return v.String()
	default:
		return v
	}
}

func RuntimeLuaConvertMapString(l *lua.LState, data map[string]string) *lua.LTable {
	lt := l.CreateTable(0, len(data))

	for k, v := range data {
		lt.RawSetString(k, RuntimeLuaConvertValue(l, v))
	}

	return lt
}

func RuntimeLuaConvertMap(l *lua.LState, data map[string]interface{}) *lua.LTable {
	lt := l.CreateTable(0, len(data))

	for k, v := range data {
		lt.RawSetString(k, RuntimeLuaConvertValue(l, v))
	}

	return lt
}

func RuntimeLuaConvertMapInt64(l *lua.LState, data map[string]int64) *lua.LTable {
	lt := l.CreateTable(0, len(data))

	for k, v := range data {
		lt.RawSetString(k, RuntimeLuaConvertValue(l, v))
	}

	return lt
}

func RuntimeLuaConvertValue(l *lua.LState, val interface{}) lua.LValue {
	if val == nil {
		return lua.LNil
	}

	// Types looked up from:
	// https://golang.org/pkg/encoding/json/#Unmarshal
	// https://developers.google.com/protocol-buffers/docs/proto3#scalar
	// More types added based on observations.
	switch v := val.(type) {
	case bool:
		return lua.LBool(v)
	case string:
		return lua.LString(v)
	case []byte:
		return lua.LString(v)
	case float32:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	case int:
		return lua.LNumber(v)
	case int8:
		return lua.LNumber(v)
	case int32:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case uint:
		return lua.LNumber(v)
	case uint8:
		return lua.LNumber(v)
	case uint32:
		return lua.LNumber(v)
	case uint64:
		return lua.LNumber(v)
	case map[string][]string:
		lt := l.CreateTable(0, len(v))
		for k, v := range v {
			lt.RawSetString(k, RuntimeLuaConvertValue(l, v))
		}
		return lt
	case map[string]string:
		return RuntimeLuaConvertMapString(l, v)
	case map[string]int64:
		return RuntimeLuaConvertMapInt64(l, v)
	case map[string]interface{}:
		return RuntimeLuaConvertMap(l, v)
	case []string:
		lt := l.CreateTable(len(val.([]string)), 0)
		for k, v := range v {
			lt.RawSetInt(k+1, lua.LString(v))
		}
		return lt
	case []interface{}:
		lt := l.CreateTable(len(val.([]interface{})), 0)
		for k, v := range v {
			lt.RawSetInt(k+1, RuntimeLuaConvertValue(l, v))
		}
		return lt
	case time.Time:
		return lua.LNumber(v.UTC().Unix())
	case nil:
		return lua.LNil
	default:
		// Never return an actual Go `nil` or it will cause nil pointer dereferences inside gopher-lua.
		return lua.LNil
	}
}
