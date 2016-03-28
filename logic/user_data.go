package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"app-server/config"
	"app-server/define"
	"app-server/models"
	"app-server/pkg/mongodb"
	"app-server/pkg/redis"
	"app-server/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

const (
	RankLimit = 100000
)

/////////////////////////////////////////////////////////
// 封装用户数据的数据库更新
/////////////////////////////////////////////////////////
func UpdateUserById(userId string, update interface{}) error {
	id := bson.M{"_id": bson.ObjectIdHex(userId)}
	err := mongodb.Update(define.AccountCollection, id, update)
	return err
}

/////////////////////////////////////////////////////////
// 用户数据库查找相关
/////////////////////////////////////////////////////////
// 根据userid获取user
func FindUserDataById(userId string) (*models.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("userid invaild")
	}

	var user models.User
	id := bson.ObjectIdHex(userId)
	err := mongodb.SelectById(define.AccountCollection, id, nil, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 根据手机号或红包号查找用户
func FindUserByAccount(username string) *models.User {
	accountType := utils.CheckAccout(username)

	var query bson.M
	user := new(models.User)
	switch accountType {
	case utils.E_AccountPhone:
		query = bson.M{"phone": username}
	case utils.E_AccountString:
		query = bson.M{"account": username}
	default:
		return nil
	}

	err := mongodb.SelectOne(define.AccountCollection, query, nil, user)
	if err != nil {
		return nil
	}

	if !user.Id_.Valid() {
		return nil
	}
	updateUserCache(user.Id_.Hex(), user)

	return user
}

// 通过第三方openid查询用户数据
func FindUserByOpenId(openId string, platform int) *models.User {
	var query bson.M
	switch platform {
	case define.E_Weixin:
		query = bson.M{"weixin_openid": openId}
	case define.E_Weibo:
		query = bson.M{"weibo_openid": openId}
	case define.E_QQ:
		query = bson.M{"qq_openid": openId}
	default:
		return nil
	}

	user := new(models.User)
	err := mongodb.SelectOne(define.AccountCollection, query, nil, user)
	if err != nil {
		return nil
	}

	if !user.Id_.Valid() {
		return nil
	}
	updateUserCache(user.Id_.Hex(), user)

	return user
}

// 根据账号查询用户id
func FindUserIdByAccount(account string) string {
	var userId struct {
		Id bson.ObjectId `bson:"_id"`
	}

	var query bson.M
	accountType := utils.CheckAccout(account)
	switch accountType {
	case utils.E_AccountPhone:
		query = bson.M{"phone": account}
	case utils.E_AccountString:
		query = bson.M{"account": account}
	default:
		return ""
	}

	selector := bson.M{"_id": 1}
	err := mongodb.SelectOne(define.AccountCollection, query, selector, &userId)
	if err != nil {
		return ""
	}

	return userId.Id.Hex()
}

/////////////////////////////////////////////////////////
// 用户数据相关操作
/////////////////////////////////////////////////////////
// 更新用户的账户余额
func UpdateMoney(userId string, money int64) {
	update := bson.M{"$inc": bson.M{"money": money}}
	err := UpdateUserById(userId, update)
	if err == nil { // 更新缓存
		go IncUserCacheField(userId, define.EUser_Money, money)
	}
}

// 更新用户收到的红包统计
func updateIncome(userId string, money int) {
	if money > 0 {
		update := bson.M{"$inc": bson.M{"income": money}}
		err := UpdateUserById(userId, update)
		if err == nil {
			go IncUserCacheField(userId, define.EUser_Income, money)

			//红包排行榜
			curIncome, gErr := getUserField(userId, define.EUser_Income)
			if gErr == nil {
				old, _ := strconv.Atoi(curIncome)
				cur := money + old
				fmt.Println("------> 最新的红包总收入 = ", cur)
				redis.Do("ZADD", define.IncomeRank, cur, userId)
				count, _ := redis.GetInt("ZCARD", define.IncomeRank)
				if count > RankLimit {
					redis.Do("ZREMRANGEBYRANK", define.IncomeRank, RankLimit, "-1")
				}
			}
		}
	}
}

// 更新发红包的统计数据
func updateOutcome(userId string, money int) {
	cacheErr := updateCacheTTL(userId)
	year := time.Now().Format("2006")
	key0 := fmt.Sprintf("outcome.%s.0", year)
	key1 := fmt.Sprintf("outcome.%s.1", year)

	update := bson.M{"$inc": bson.M{key0: money, key1: 1}}
	err := UpdateUserById(userId, update)
	if err == nil && cacheErr == nil { // 更新缓存
		args := []interface{}{
			define.UserCachePrefix + userId,   //KEYS
			define.EUser_Outcome, year, money, // ARGV
		}
		total, luaErr := redis.DoLuaInt(define.GUser_Outcome, 1, args...)
		if luaErr == nil && total > 0 {
			// 老板排行榜
			go func() {
				fmt.Println("发出的红包 = ", total)
				redis.Do("ZADD", define.OutcomeRank, total, userId)
				count, _ := redis.GetInt("ZCARD", define.OutcomeRank)
				if count > RankLimit {
					redis.Do("ZREMRANGEBYRANK", define.OutcomeRank, RankLimit, "-1")
				}
			}()
		}
	}
}

// 查询用户发出的红包统计数据
func getOutComeData(userId string) map[string][]int {
	str, err := getUserField(userId, define.EUser_Outcome)
	if err != nil {
		return nil
	}

	var outcome map[string][]int
	err = json.Unmarshal([]byte(str), &outcome)
	if err != nil {
		return nil
	}

	return outcome
}

// 获取总收入排名
func GetIncomeRank(userId string) int {
	rank, err := redis.GetInt("ZRANK", define.IncomeRank, userId)
	if err != nil { // 超过十万
		return 0
	}

	return rank
}

// 获取每日已抢红包的次数
func GetDailyGrab(userId string) int {
	str, err := getUserField(userId, define.EUser_DailyGrab)
	if err != nil {
		return 0
	}

	var list []int64
	err = json.Unmarshal([]byte(str), &list)
	if err != nil {
		return 0
	}

	if len(list) == 2 {
		date := time.Now().Format("20060102")
		last := time.Unix(list[1], 0).Format("20060102")
		if date != last {
			return 0
		} else {
			return int(list[0])
		}
	}
	return 0
}

// 更新每日抢红包的次数和最后一次抢红包的时间
func UpdateDailyGrab(userId string) {
	old := GetDailyGrab(userId)
	daily := []int64{int64(old + 1), time.Now().Unix()}
	update := bson.M{"$set": bson.M{"daily_grab": daily}}
	err := UpdateUserById(userId, update)
	if err == nil {
		bytes, _ := json.Marshal(daily)
		UpdateUserCacheField(userId, define.EUser_DailyGrab, string(bytes))
	}
}

// 根据积分配置获取用户等级
func GetUserLevel(userId string) int {
	v, err := getUserField(userId, define.EUser_Point)
	if err != nil {
		return -1
	}

	point, _ := strconv.Atoi(v)
	return config.GetUserLevelByPoint(point)
}

// 更新用户积分
func updateUserPoint(userId string, point int) {
	update := bson.M{"$inc": bson.M{"point": point}}
	err := UpdateUserById(userId, update)
	if err == nil {
		IncUserCacheField(userId, define.EUser_Point, point)

		// 更新等级排行榜
		go func(userId string) {
			lv := GetUserLevel(userId)
			redis.Do("ZADD", define.LevelRank, lv, userId)
			count, _ := redis.GetInt("ZCARD", define.IncomeRank)
			if count > RankLimit {
				redis.Do("ZREMRANGEBYRANK", define.IncomeRank, RankLimit, "-1")
			}
		}(userId)
	}
}

// 根据红包类型获取每日每种类型红包的抢夺限制
// 目前只有分享类和公众号类有限制
func getGrabLimitByType(userId string, redpktType int) int {
	var feild string
	switch redpktType {
	case define.ERedpkt_Share:
		feild = define.EUser_ShareLimit
	case define.ERedpkt_OA:
		feild = define.EUser_OALimit
	default:
		return -1
	}

	str, err := getUserField(userId, feild)
	if err != nil {
		return -1
	}

	var list []int64
	err = json.Unmarshal([]byte(str), &list)
	if err != nil {
		return -1
	}

	if len(list) == 2 {
		date := time.Now().Format("20060102")
		last := time.Unix(list[1], 0).Format("20060102")
		if date != last {
			return 0
		} else {
			return int(list[0])
		}
	}
	return 0
}

func updateGrabLimit(userId string, redpktType int) {
	used := getGrabLimitByType(userId, redpktType)
	if used < 0 {
		return
	}

	var feild string
	var update bson.M
	limit := []int64{int64(used + 1), time.Now().Unix()}
	switch redpktType {
	case define.ERedpkt_Share:
		feild = define.EUser_ShareLimit
		update = bson.M{"$set": bson.M{"share_limit": limit}}
	case define.ERedpkt_OA:
		feild = define.EUser_OALimit
		update = bson.M{"$set": bson.M{"oa_limit": limit}}
	default:
		return
	}

	err := UpdateUserById(userId, update)
	if err == nil {
		bytes, _ := json.Marshal(limit)
		UpdateUserCacheField(userId, feild, string(bytes))
	}
}

// 更新用户的登陆次数
func updateLoginCount(userData *models.User) {
	userId := userData.Id_.Hex()
	update := bson.M{"$inc": bson.M{"login_count": 1}}
	err := UpdateUserById(userId, update)
	if err == nil {
		IncUserCacheField(userId, define.EUser_LoginCount, 1)

		// 登陆积分
		cfg := config.GetPointCfg(define.EPoint_Login)
		if cfg != nil {
			t := userData.LoginCount + 1
			if t == cfg.Limit {
				updateUserPoint(userId, cfg.Point)
			}
		}
	}
}
