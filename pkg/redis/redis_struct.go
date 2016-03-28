// 封装redis的struct操作，主要针对hash表操作优化

package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	goredis "github.com/garyburd/redigo/redis"
)

// 结构体字段
type fieldSpec struct {
	fieldName string // 结构体字段名
	name      string // 结构体字段名tag
	omitEmpty bool
	inline    bool
}

// 反射结构体
func compileStruct(t reflect.Type) map[string]*fieldSpec {
	fields := make(map[string]*fieldSpec)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fs := &fieldSpec{name: f.Name, fieldName: f.Name}

		tag := f.Tag.Get("redis")
		p := strings.Split(tag, ",")
		if len(p) > 0 {
			if p[0] == "-" {
				continue
			}

			// 结构体字段名
			if len(p[0]) > 0 {
				fs.name = p[0]
			}

			for _, s := range p[1:] {
				switch s {
				case "omitempty":
					fs.omitEmpty = true
				case "inline":
					fs.inline = true
				default:
					panic(fmt.Errorf("redigo: unknown field tag %s for type %s", s, t.Name()))
				}
			}
		}

		fields[fs.name] = fs
	}

	return fields
}

// 拆封结构体
func flattenStruct(args goredis.Args, v reflect.Value, inlineFlag bool) goredis.Args {
	fields := compileStruct(v.Type())
	for name, spec := range fields {
		fv := v.FieldByName(spec.fieldName)
		if spec.omitEmpty {
			var empty = false
			switch fv.Kind() {
			case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
				empty = fv.Len() == 0
			case reflect.Bool:
				empty = !fv.Bool()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				empty = fv.Int() == 0
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				empty = fv.Uint() == 0
			case reflect.Float32, reflect.Float64:
				empty = fv.Float() == 0
			case reflect.Interface, reflect.Ptr:
				empty = fv.IsNil()
			}
			if empty {
				continue
			}
		}

		kind := fv.Kind()
		if kind == reflect.Ptr {
			kind = fv.Elem().Kind()
		}
		if kind == reflect.Map || kind == reflect.Slice || kind == reflect.Struct {
			if spec.inline && inlineFlag && kind == reflect.Struct { // 第一层才能打散
				args = flattenStruct(args, fv, false)
			} else {
				out, err := json.Marshal(fv.Interface())
				if err == nil {
					args = append(args, name, string(out))
				}
			}
		} else {
			args = append(args, name, fv.Interface())
		}
	}

	return args
}

// 反序列化字段
// inlineFlag = false 表示使inline tag失效，一般用于嵌套的结构体，因为不支持结构体内部的子结构体使用inline
func setField(rv reflect.Value, src map[string][]byte, inlineFlag bool) error {
	var err error
	rt := rv.Type()
	fields := compileStruct(rt)
	for _, fs := range fields {
		d := rv.FieldByName(fs.fieldName)
		if fs.inline && inlineFlag {
			typ := d.Type()
			res := reflect.New(typ).Interface()
			rvInline := reflect.ValueOf(res)
			if rvInline.Kind() != reflect.Ptr || rvInline.IsNil() {
				return errors.New("[Error] setField it's must be prt……")
			}

			rvInline = rvInline.Elem()
			if rvInline.Kind() != reflect.Struct {
				return errors.New("[Error] setField it's must be struct……")
			}

			setField(rvInline, src, false) // 第二层不支持inline
			d.Set(reflect.ValueOf(res).Elem())
		} else {
			s, ok := src[fs.name]
			if !ok || len(s) == 0 {
				continue
			}

			if d.CanSet() {
				switch d.Type().Kind() {
				case reflect.Float32, reflect.Float64:
					var x float64
					x, err = strconv.ParseFloat(string(s), d.Type().Bits())
					d.SetFloat(x)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					var x int64
					x, err = strconv.ParseInt(string(s), 10, d.Type().Bits())
					d.SetInt(x)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					var x uint64
					x, err = strconv.ParseUint(string(s), 10, d.Type().Bits())
					d.SetUint(x)
				case reflect.Bool:
					var x bool
					x, err = strconv.ParseBool(string(s))
					d.SetBool(x)
				case reflect.String:
					d.SetString(string(s))
				case reflect.Slice:
					typ := d.Type()
					res := reflect.New(typ).Interface()
					e := json.Unmarshal(s, res)
					if e == nil {
						d.Set(reflect.ValueOf(res).Elem())
					}
				case reflect.Map:
					typ := d.Type()
					res := reflect.New(typ).Interface()
					e := json.Unmarshal(s, res)
					if e == nil {
						d.Set(reflect.ValueOf(res).Elem())
					}
				case reflect.Struct:
					typ := d.Type()
					res := reflect.New(typ).Interface()
					e := json.Unmarshal(s, res)
					if e == nil {
						d.Set(reflect.ValueOf(res).Elem())
					}
				case reflect.Ptr:
					typ := d.Type().Elem()
					res := reflect.New(typ).Interface()
					e := json.Unmarshal(s, res)
					if e == nil {
						d.Set(reflect.ValueOf(res))
					}
				default:
					continue
				}
			}
		}
	}

	return err
}

/////////////////////////////////////////////////////////
// Hash 结构体操作
/////////////////////////////////////////////////////////
// 支持多k-v操作，需要指定key
// 例如：MSET,HMSET等
// arg字段支持slice、map和struct
func MSet(cmdName string, key string, arg interface{}) (interface{}, error) {
	return Do(cmdName, goredis.Args{}.Add(key).AddFlat(arg)...)
}

// hash表的结构体set
func StructHMset(key string, arg interface{}) (interface{}, error) {
	args := goredis.Args{}.Add(key)

	rv := reflect.ValueOf(arg)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("[Error] it's must be prt……")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("[Error]StructHMset it's must be struct……")
	}

	args = flattenStruct(args, rv, true)
	return Do("HMSET", args...)
}

func HGetall(key string, args interface{}) error {
	rv := reflect.ValueOf(args)
	if rv.Kind() != reflect.Ptr { //reply必须是指针
		return errors.New("[Error] it's must be prt……")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("[Error] HGetall it's must be struct……")
	}

	// get 哈希表所有元素
	src, herr := GetValues("HGETALL", key)
	if herr != nil {
		return herr
	}

	if len(src) == 0 {
		return errors.New("empty hash")
	}

	if len(src)%2 != 0 {
		return errors.New("redigo.ScanStruct: number of values not a multiple of 2")
	}

	maps := make(map[string][]byte)
	for i := 0; i < len(src); i += 2 {
		k := src[i+1]
		if k == nil {
			continue
		}
		name, ok := src[i].([]byte)
		if !ok {
			return errors.New("hgetall struct field name error")
		}

		s := k.([]byte)
		maps[string(name)] = s
	}

	err := setField(rv, maps, true)
	return err
}
