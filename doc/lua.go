package doc

import (
	"errors"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"datatool.deflinhec.dev/doc/typecast"
	lua "github.com/yuin/gopher-lua"
)

func luaState() *lua.LState {
	l := lua.NewState(lua.Options{
		CallStackSize:       128,
		RegistryGrowStep:    128,
		RegistryMaxSize:     128,
		SkipOpenLibs:        true,
		IncludeGoStackTrace: true,
	})
	for name, lib := range map[string]lua.LGFunction{
		lua.BaseLibName:   lua.OpenBase,
		lua.TabLibName:    lua.OpenTable,
		lua.IoLibName:     lua.OpenIo,
		lua.StringLibName: lua.OpenString,
		lua.MathLibName:   lua.OpenMath,
		lua.LoadLibName:   lua.OpenPackage,
	} {
		l.Push(l.NewFunction(lib))
		l.Push(lua.LString(name))
		l.Call(1, 0)
	}
	return l
}

func luaLoadEntries(l *lua.LState) ([]string, error) {
	err := l.DoString(`
	local function expand(t, paths)
		paths = paths or {}
		local parent = paths[#paths] or ""
		for k, v in pairs(t) do
			local skip = false
			if type(k) == 'string' then
				if k:sub(1,1) == '_' then
					skip = true
				end
			end
			if not skip then
				if type(v) == 'table' then
					local path = parent
					if type(k) == 'number' then
						path = path .. '['..k..']'
					elseif #path > 0 then
						path = path .. '.' .. k
					else
						path = path .. k
					end
					table.insert(paths, path)
					expand(v, paths)
				end
			end
		end
		return paths
	end
	return expand(package.loaded)
	`)
	if err != nil {
		return []string{}, err
	}
	defer l.Pop(l.GetTop())
	value := typecast.RuntimeLuaConvertLuaValue(l.Get(-1))
	if value, ok := value.([]interface{}); ok {
		entries := make([]string, 0)
		for _, entry := range value {
			path := fmt.Sprint(entry)
			if !strings.Contains(path, ".") {
				continue
			} else if strings.Contains(path, ".md5sum") {
				continue
			}
			entries = append(entries, path)
		}
		return entries, nil
	}
	return []string{}, errors.New(translator.Sprintf("無法載入數據 %T", value))
}

func luaLoadValue(l *lua.LState, entry string) (interface{}, error) {
	err := l.DoString(`
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
	local function cast(value)
		if type(value) == 'table' then
			if is_array(value) then
				return value
			end
			local table = {}
			for k, v in pairs(value) do
				if type(k) == 'number' then
					k = '#'..k
				end
				table[k] = cast(v)
			end
			return table
		end
		return value
	end
	return cast(` + entry + `)
	`)
	if err != nil {
		return nil, err
	}
	defer l.Pop(l.GetTop())
	return typecast.RuntimeLuaConvertLuaValue(l.Get(-1)), nil
}

type luaFile struct {
	fileInfo
}

func (f *luaFile) Path() string {
	return filepath.Join(f.dir, f.name)
}

func (f *luaFile) Module() string {
	return f.module
}

func (f *luaFile) Field() string {
	return f.field
}

func (f *luaFile) Md5Sum() string {
	return f.md5sum
}

func (f *luaFile) read() (interface{}, error) {
	l := luaState()
	defer l.Close()
	fs, err := ioutil.ReadFile(filepath.Join(f.dir, f.name))
	if err != nil {
		return nil, err
	}
	fn, err := l.LoadString(string(fs))
	if err != nil {
		return nil, err
	}
	l.Push(fn)
	err = l.PCall(0, 0, nil)
	if err != nil {
		return nil, err
	}
	entries, err := luaLoadEntries(l)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, errors.New(translator.Sprintf("沒有數據"))
	}
	value, err := luaLoadValue(l, entries[0])
	if err != nil {
		return nil, err
	}
	field := strings.Split(entries[0], ".")
	f.module, f.field, f.md5sum = field[0], field[1], md5Sum(value)
	return value, nil
}

func maxInt(value, max int) int {
	return int(math.Max(float64(value), float64(max)))
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func luaArray(values []interface{}) bool {
	for _, value := range values {
		if value == nil {
			return false
		}
	}
	return true
}

func luaFormat(value interface{}, depth int, trait string) string {
	var ctx string
	switch value := value.(type) {
	case int:
		ctx += fmt.Sprint(value)
	case int16:
		ctx += fmt.Sprint(value)
	case int32:
		ctx += fmt.Sprint(value)
	case int64:
		ctx += fmt.Sprint(value)
	case uint:
		ctx += fmt.Sprint(value)
	case uint16:
		ctx += fmt.Sprint(value)
	case uint32:
		ctx += fmt.Sprint(value)
	case uint64:
		ctx += fmt.Sprint(value)
	case string:
		ctx += strconv.QuoteToGraphic(value)
	case float32:
		ctx += fmt.Sprint(value)
	case float64:
		ctx += fmt.Sprint(value)
	case bool:
		ctx += fmt.Sprint(value)
	case []interface{}:
		if !luaArray(value) {
			values := make(map[string]interface{})
			for i, v := range value {
				if v != nil {
					values[fmt.Sprintf("#%v", i+1)] = v
				}
			}
			return luaFormat(values, depth, trait)
		}
		var brackets int
		lhs := make([]string, 0, len(value))
		rhs := make([]string, 0, len(value))
		for i := 0; i < len(value); i++ {
			var l, r, e string
			if len(value) >= 25550 {
				l = fmt.Sprintf("[%v]=", i+1)
			}
			e = fmt.Sprintf("%v[%v]", trait, i+1)
			r = luaFormat(value[i], depth+4, e)
			lhs, rhs = append(lhs, l), append(rhs, r)
			brackets += strings.Count(r, "{")
		}
		var line, chunk string
		for i := 0; i < len(value); i++ {
			line += lhs[i] + rhs[i] + ","
			if len(line) > 60 {
				chunk += strings.Repeat(" ", depth+4)
				chunk += line + "\n"
				line = ""
			}
		}
		if len(line) > 0 {
			chunk += strings.Repeat(" ", depth+4)
			chunk += line + "\n"
			line = ""
		}
		ctx += "{"
		if brackets == 0 {
			ctx += fmt.Sprintf(" -- total: %v", len(value))
			if strings.Count(chunk, "\n") > 1 {
				values, l := make([]string, 0, len(value)), 0
				sections := strings.Split(chunk, ",")
				for i := 0; i < len(value); i++ {
					value := sections[i]
					value = strings.Trim(value, "\n")
					value = strings.Trim(value, " ")
					values = append(values, value)
				}
				for i := 0; i < len(value); i++ {
					l = maxInt(l, len(values[i]))
				}
				line, chunk = strings.Repeat(" ", depth+4), ""
				format := fmt.Sprintf("%%%vs", l)
				for i := 0; i < len(value); i++ {
					line += fmt.Sprintf(format, values[i]) + ","
					if len(line) > 60 {
						chunk += line + "\n"
						line = strings.Repeat(" ", depth+4)
					}
				}
				if len(strings.Trim(line, " ")) > 0 {
					chunk += line + "\n"
				}
			}
		}
		ctx += "\n"
		ctx += chunk
		ctx += strings.Repeat(" ", depth)
		ctx += "}"
	case map[string]interface{}:
		ctx += "{" + fmt.Sprintf(" -- %v", trait) + "\n"
		ctx += strings.Repeat(" ", depth+4)
		keys := make(map[uint64]string)
		idxs := make([]uint64, 0, len(value))
		for k := range value {
			idx := uint64(hash(k))
			if strings.HasPrefix(k, "#") {
				idx, _ = strconv.ParseUint(k[1:], 10, 32)
			}
			idxs = append(idxs, idx)
			keys[idx] = k
		}
		sort.Slice(idxs, func(i, j int) bool { return idxs[i] < idxs[j] })
		for _, i := range idxs {
			k := keys[i]
			var e string
			var keypair string
			if strings.HasPrefix(k, "#") {
				e = fmt.Sprintf("%v[%v]", trait, k[1:])
				keypair += fmt.Sprintf("[%v]=", k[1:])
			} else if k == "true" || k == "false" {
				e = fmt.Sprintf("%v[\"%v\"]", trait, k)
				keypair += fmt.Sprintf("[\"%v\"]=", k)
			} else {
				e = fmt.Sprintf("%v.%v", trait, k)
				keypair += fmt.Sprintf("%v=", k)
			}
			keypair += luaFormat(value[k], depth+4, e) + ","
			keypair += "\n" + strings.Repeat(" ", depth+4)
			ctx += keypair
		}
		if len(keys) > 0 {
			ctx = ctx[:len(ctx)-4]
		}
		ctx += "}"
	}
	return ctx
}

func (f *luaFile) write(value interface{}) error {
	fs, err := os.Create(filepath.Join(f.dir, f.name))
	if err != nil {
		return err
	}
	if len(f.module) == 0 {
		return errors.New("invalid module")
	} else if len(f.field) == 0 {
		return errors.New("invalid field")
	}
	defer fs.Close()
	f.md5sum = md5Sum(value)
	fs.WriteString("-- $id$\n")
	fs.WriteString(`module("` + f.module + `")` + "\n")
	fs.WriteString("md5sum = md5sum or {}\n")
	fs.WriteString(`md5sum.` + f.field + `=`)
	fs.WriteString(`"` + f.md5sum + `"` + "\n")
	fs.WriteString(f.field + "=\n")
	route := fmt.Sprintf("%v.%v", f.module, f.field)
	fs.WriteString(luaFormat(value, 0, route))
	return nil
}
