// redis lua脚本

package redis

import (
	goredis "github.com/garyburd/redigo/redis"
)

// 注意：使用script传入的key 和 args 都是string
func DoLua(luaStr string, keyCount int, keyAndArgs ...interface{}) (interface{}, error) {
	c := GetRedisConn()
	script := goredis.NewScript(keyCount, luaStr)
	reply, err := script.Do(c, keyAndArgs...)
	return reply, err
}

func DoLuaUint(luaStr string, keyCount int, keyAndArgs ...interface{}) (uint32, error) {
	reply, err := goredis.Uint64(DoLua(luaStr, keyCount, keyAndArgs...))
	return uint32(reply), err
}

func DoLuaInt(luaStr string, keyCount int, keyAndArgs ...interface{}) (int, error) {
	return goredis.Int(DoLua(luaStr, keyCount, keyAndArgs...))
}

func DoLuaInt64(luaStr string, keyCount int, keyAndArgs ...interface{}) (int64, error) {
	return goredis.Int64(DoLua(luaStr, keyCount, keyAndArgs...))
}
