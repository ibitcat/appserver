// 封装的redis的get操作

package redis

import (
	"errors"
	"log"
	"reflect"

	goredis "github.com/garyburd/redigo/redis"
)

// 封装redigo的get方法
func GetBool(cmdName string, args ...interface{}) (bool, error) {
	return goredis.Bool(Do(cmdName, args...))
}

func GetInt64(cmdName string, args ...interface{}) (int64, error) {
	return goredis.Int64(Do(cmdName, args...))
}

func GetInt(cmdName string, args ...interface{}) (int, error) {
	return goredis.Int(Do(cmdName, args...))
}

func GetBytes(cmdName string, args ...interface{}) ([]byte, error) {
	return goredis.Bytes(Do(cmdName, args...))
}

func GetString(cmdName string, args ...interface{}) (string, error) {
	return goredis.String(Do(cmdName, args...))
}

func GetStringMap(cmdName string, args ...interface{}) (map[string]string, error) {
	return goredis.StringMap(Do(cmdName, args...))
}

func GetStrings(cmdName string, args ...interface{}) ([]string, error) {
	return goredis.Strings(Do(cmdName, args...))
}

func GetValues(cmdName string, args ...interface{}) ([]interface{}, error) {
	return goredis.Values(Do(cmdName, args...))
}

func ScanSlice(cmdName string, args ...interface{}) error {
	l := len(args)
	if l < 2 {
		return errors.New("[Error]get slice param fail……")
	}

	values, err := goredis.Values(Do(cmdName, args[:l-1]...))
	goredis.ScanSlice(values, args[l])
	return err
}

func ScanStruct(cmdName string, args ...interface{}) error {
	// rv := reflect.ValueOf(reply)
	// if rv.Kind() != reflect.Ptr { //reply必须是指针
	// 	return errors.New("[Error] it's must be prt……")
	// }
	return nil
}

// 指定返回值的get操作
// reply 为返回值,需要传入指针（必须是指针）
// 该方法主要封装了redigo内置的类型转换方法
func ReplyGet(cmdName string, reply interface{}, args ...interface{}) error {
	rv := reflect.ValueOf(reply)
	if rv.Kind() != reflect.Ptr { //reply必须是指针
		return errors.New("[Error] it's must be prt……")
	}

	r, err := Do(cmdName, args...)
	if rv.Type().Elem().Kind() == reflect.Struct { //如果返回值是结构体指针
		v, vErr := goredis.Values(r, err)
		if vErr != nil {
			return vErr
		}

		//封装redis的ScanStruct方法
		log.Println("ReplyGet--------struct")
		e := goredis.ScanStruct(v, reply)
		return e
	} else {
		switch result := reply.(type) {
		case *[][]byte:
			out, outErr := goredis.ByteSlices(r, err)
			if outErr != nil {
				return outErr
			}
			*result = out
			//copy(result, out)
		case *[]byte:
			out, outErr := goredis.Bytes(r, err)
			if outErr != nil {
				return outErr
			}
			*result = out
			//copy(result, out)
		case *int64:
			out, outErr := goredis.Int64(r, err)
			if outErr != nil {
				return outErr
			}
			*result = out
		case *string:
			out, outErr := goredis.String(r, err)
			if outErr != nil {
				return outErr
			}
			*result = out
		case *map[string]string:
			out, outErr := goredis.StringMap(r, err)
			if outErr != nil {
				return outErr
			}
			*result = out
		default:
			return errors.New("[Error]redis reply parse fail……")
		}
	}

	if err != nil {
		return err
	}

	if r == nil {
		return goredis.ErrNil
	}

	return nil
}
