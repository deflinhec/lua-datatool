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
	lua "github.com/yuin/gopher-lua"
)

const (
	KeyEscapeRune = '#'
)

// Since Lua allow non-string key within table, we need to escape them
// by prepending a key escape rune in order to cast as map[string]interface{}.
// Key escaoe rune '#' indicate the key which is orignal an number key,
// EscapeNonStringKeys escapes non-string keys of the table.
func EscapeNonStringKeys(l *lua.LState, value lua.LValue) lua.LValue {
	if err := l.DoString(__helper_script__); err != nil {
		panic(err)
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("__escape__"),
		NRet:    1,
		Protect: true,
	}, value); err != nil {
		panic(err)
	}
	return l.Get(-1)
}

// UnescapeNonStringKeys unescapes non-string keys of the table.
func UnescapeNonStringKeys(l *lua.LState, value lua.LValue) lua.LValue {
	if err := l.DoString(__helper_script__); err != nil {
		panic(err)
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("__unescape__"),
		NRet:    1,
		Protect: true,
	}, value); err != nil {
		panic(err)
	}
	return l.Get(-1)
}

// This is a helper function for non-string keys.
const __helper_script__ = `
local __rune__ = '` + string(KeyEscapeRune) + `'
local function is_array(t)
	local i = 0
	for _ in pairs(t) do
		i = i + 1
		if t[i] == nil then
			return false
		end
	end
	return true
end
function __escape__(value)
	if type(value) == 'table' then
		if is_array(value) then
			return value
		end
		local table = {}
		for k, v in pairs(value) do
			if type(k) == 'number' then
				k = __rune__..k
			end
			table[k] = __escape__(v)
		end
		return table
	end
	return value
end
function __unescape__(value)
	if type(value) == 'table' then
		if is_array(value) then
			return value
		end
		local table = {}
		for k, v in pairs(value) do
			if type(k) == 'string' then
				if k:sub(1,1) == __rune__ then
					k = assert(tonumber(k:sub(2)))
				end
			end
			table[k] = __unescape__(v)
		end
		return table
	end
	return value
end
`
