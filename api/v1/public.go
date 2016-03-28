/*
app后端api
*/

package v1

import (
	//"encoding/json"
	"fmt"
	//"strconv"
	"time"

	"app-server/models"
	//"app-server/pkg/middleware"
	//"app-server/logic"
	"app-server/pkg/redis"
	//"app-server/pkg/sdk/jpush"

	"github.com/gin-gonic/gin"
)

// @SubApi 公共接口 [/pub]
func Public(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/pub")

	//route.Use(middleware.JwtAuthMiddleware())
	route.GET("/debug", debug)
	route.GET("/servertime", getServerTime)
	route.GET("/test", publiceTest)
	route.POST("/posttest", publicePostTest)
}

// @Title getServerTime
// @Description 获取服务器时间
// @Accept  json
// @Success 200 {string} string "获取服务器时间成功"
// @Resource /pub
// @Router /pub/servertime [get]
func getServerTime(c *gin.Context) {
	c.JSON(200, gin.H{"time": time.Now().Unix()})
}

// @Title publiceTest
// @Description 公共接口测试
// @Accept  json
// @Success 200 {string} string "测试成功"
// @Failure 401 {object} models.APIError "测试失败"
// @Resource /pub
// @Router /pub/test [get]
func publiceTest(c *gin.Context) {
	fmt.Println("===> url ", c.Request.Referer())
	startTime := time.Now()

	type F struct {
		A string `json:"a"`
	}

	var p1, p2 struct {
		Title   string            `redis:"title"`
		Author  string            `redis:"author"`
		Body    string            `redis:"body"`
		Test    []string          `redis:"test"`
		Maptest map[string]string `redis:"map"`
		FF      *F                `redis:"f"`
	}

	p1.Author = "domi"
	p1.Title = "标题"
	p1.Body = "内容"
	p1.Test = []string{"test1", "test2"}
	p1.Maptest = map[string]string{"m1": "map1", "m2": "map2"}
	p1.FF = &F{"aaaaa"}
	fmt.Println(p1)
	fmt.Println(redis.StructHMset("hash", &p1))

	fmt.Println("---------set-------\n")
	redis.HGetall("hash", &p2)
	fmt.Println("p2=", p2, p2.FF)

	fmt.Println("--------get--------\n")
	var mapString map[string]string
	err := redis.ReplyGet("HGETALL", &mapString, "hash")
	fmt.Println("ReplyGet = ", mapString, err)

	// var unused int = 0
	// reply, _ := redis.Do("LRANGE", "list", 0, -1)
	// for _, vv := range reply.([]interface{}) { // 回收未发完的红包
	// 	money, _ := strconv.Atoi(string(vv.([]byte)))
	// 	unused += money
	// }

	//fmt.Println("总金额 = ", unused)
	fmt.Println("消耗时间 time = ", time.Now().Sub(startTime))
	c.JSON(200, models.APIError{200, "test ok"})
}

// @Title postTest
// @Description 公共接口post测试
// @Accept  json
// @Success 200 {string} string "测试成功"
// @Failure 401 {string} string "测试失败"
// @Resource /pub
// @Router /pub/posttest [post]
func publicePostTest(c *gin.Context) {
	fmt.Println("body = ", c.Request.Body)
	c.JSON(200, gin.H{"key": "test", "payload": gin.H{"success": true, "name": "test"}})
}

// @Title debug
// @Description 调试接口
// @Accept  json
// @Success 200 {string} string "测试成功"
// @Failure 401 {string} string "测试失败"
// @Resource /pub
// @Router /pub/debug [get]
func debug(c *gin.Context) {
	/*lua := `
		local a = [[{"age":12,"name":"domi"}]]
		local date = cjson.decode(a)
		for k,v in pairs(date) do
			redis.log(redis.LOG_NOTICE,"cjson---->",k,v,type(k),type(v))
		end

		local b = {1,2,"test"}
		redis.log(redis.LOG_NOTICE,"encode ---> ", cjson.encode(b))
	`
	stc := struct {
		Name string
		Age  int
	}{"domi", 21}
	args := []interface{}{[]int{1, 2, 3}, "test", map[int]string{1: "a"}, stc}

	r, e := redis.DoLua(lua, 0, args...)
	fmt.Println("redis lua = ", r, e)
	*/
	//logic.GetStatistics("56838b11a5129529d0000003")
	//jpush.PushNotice()
	c.String(200, "test ok")
}
