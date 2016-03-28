/*
redis 连接池实现，参考beego cache
*/

package redis

import (
	"encoding/json"
	"errors"
	"log"

	"time"

	goredis "github.com/garyburd/redigo/redis"
)

// redis cache
type RedisCache struct {
	p        *goredis.Pool // redis 连接池
	dbNum    int           // dbnum
	conninfo string        // 连接地址
	password string        // redis管理员密码
}

var (
	redisSession *RedisCache = nil
)

// 初始化redis连接池
func NewRedisCache() {
	if redisSession == nil {
		redisSession = &RedisCache{dbNum: 0}

		err := redisSession.StartAndGC(`{"conn":"192.168.1.112:6379","password":"casablanca"}`)
		if err != nil {
			log.Println("dial redis fail, err = ", err)
			panic(err)
		}
	}
}

func Connect() *RedisCache {
	defer func() {
		if err := recover(); err != nil {
			log.Println("redis connect fail", err)
		}
	}()

	if redisSession == nil {
		NewRedisCache()
	}

	return redisSession
}

func Do(cmdName string, args ...interface{}) (interface{}, error) {
	conn := Connect()
	if conn == nil {
		return nil, errors.New("[Error]redis connect fail……")
	}
	return conn.Do(cmdName, args...)
}

func GetRedisConn() goredis.Conn {
	rc := Connect()
	if rc == nil {
		panic("[Panic]redis connect fail……")
	}
	return rc.p.Get()
}

/////////////////////////////////////////////////////////
// actually do the redis cmds
func (rc *RedisCache) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// GetMulti get cache from redis.
func (rc *RedisCache) GetMulti(keys []string) []interface{} {
	size := len(keys)
	var rv []interface{}
	c := rc.p.Get()
	defer c.Close()
	var err error
	for _, key := range keys {
		err = c.Send("GET", key)
		if err != nil {
			goto ERROR
		}
	}
	if err = c.Flush(); err != nil {
		goto ERROR
	}
	for i := 0; i < size; i++ {
		if v, err := c.Receive(); err == nil {
			rv = append(rv, v.([]byte))
		} else {
			rv = append(rv, err)
		}
	}
	return rv
ERROR:
	rv = rv[0:0]
	for i := 0; i < size; i++ {
		rv = append(rv, nil)
	}

	return rv
}

// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *RedisCache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}

	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}

	rc.conninfo = cf["conn"]
	rc.password = cf["password"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (rc *RedisCache) connectInit() {
	dialFunc := func() (c goredis.Conn, err error) {
		c, err = goredis.Dial("tcp", rc.conninfo)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}

		return
	}
	// initialize a new pool
	rc.p = &goredis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}
