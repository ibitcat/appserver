package logic

import (
	"errors"
	"fmt"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/redis"
)

/////////////////////////////////////////////////////////
// 用户缓存相关
/////////////////////////////////////////////////////////
// 更新redis缓存,默认缓存30分钟
func updateUserCache(userId string, userCache *models.User) {
	if len(userId) > 0 {
		key := define.UserCachePrefix + userId
		_, err := redis.StructHMset(key, userCache)
		if err == nil {
			redis.Do("EXPIRE", key, 1800)
		}
	}
}

// 增加用户缓存的到期时间
func updateCacheTTL(userId string) error {
	key := define.UserCachePrefix + userId
	ok, _ := redis.GetInt("EXPIRE", key, 1800)
	if ok == 0 { //key不存在了
		userInfo, err := FindUserDataById(userId)
		if err != nil {
			return err
		}
		updateUserCache(userId, userInfo)
	}
	return nil
}

// 获取用户缓存内的字段,统一返回string
func getUserField(userId, field string) (string, error) {
	key := define.UserCachePrefix + userId
	isExist, _ := redis.GetInt("EXISTS", key)
	if isExist == 0 { // 缓存过期了，重新查询数据库
		userInfo, err := FindUserDataById(userId)
		if err == nil {
			updateUserCache(userId, userInfo)
		}
	}

	value, err := redis.GetString("HGET", key, field)
	return value, err
}

// 获取用户缓存内多个字段的值,统一返回[]string
func getUserFields(userId string, field ...string) ([]string, error) {
	key := define.UserCachePrefix + userId
	isExist, _ := redis.GetInt("EXISTS", key)
	if isExist == 0 { // 缓存过期了，重新查询数据库
		userInfo, err := FindUserDataById(userId)
		if err == nil {
			updateUserCache(userId, userInfo)
		}
	}

	args := make([]interface{}, 0, len(field)+1)
	args = append(args, key)
	for _, v := range field {
		args = append(args, v)
	}

	values, err := redis.GetStrings("HMGET", args...)
	return values, err
}

// 从redis中获取user缓存
func GetUserData(userId string) (*models.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("[Error]user id is invaild")
	}
	userData := models.User{}
	Key := define.UserCachePrefix + userId
	redisErr := redis.HGetall(Key, &userData)
	fmt.Println("GetUserData -------> err = ", redisErr)
	if redisErr == nil { // key不是字符串类型或key为nil
		return &userData, nil
	}

	//重新查询数据库
	userInfo, err := FindUserDataById(userId)
	fmt.Println("FindUserDataById -------> err = ", err)
	if err == nil { // 插入到缓存
		go updateUserCache(userId, userInfo)
	}

	return userInfo, err
}

// 更新用户缓存字段(重新设置)
func UpdateUserCacheField(userId string, args ...interface{}) {
	if len(args) == 0 || len(args)%2 != 0 {
		fmt.Println("[Error] Cache 参数错误……")
		return
	}

	key := define.UserCachePrefix + userId
	isExist, _ := redis.GetInt("EXISTS", key)
	if isExist == 1 {
		newargs := make([]interface{}, 0, len(args)+1)
		newargs = append(newargs, key)
		newargs = append(newargs, args...)
		redis.Do("HMSET", newargs...)
	}
}

// 慎用
func IncUserCacheField(userId, field string, value interface{}) {
	key := define.UserCachePrefix + userId
	isExist, _ := redis.GetInt("EXISTS", key)
	if isExist == 1 {
		redis.Do("HINCRBY", key, field, value)
	}
}
