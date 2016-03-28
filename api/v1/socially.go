/*
	社交
*/

package v1

import (
	"fmt"
	"strconv"

	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 社交 [/socially]
func Socially(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/socially")

	// 不需要登陆的操作
	router.Use(middleware.JwtAuthMiddleware())

	// 融云sdk
	router.GET("/rctoken", getRongCloudToken)

	// 好友操作
	router.GET("/friend/list", getFriendList)
	router.GET("/friend/info", getFriendInfo)
	router.POST("/friend/add", addFriendById)
	router.POST("/friend/addbysearch", addFriendByAccount)
	router.POST("/friend/agree", agreeFriend)
	router.POST("/friend/refuse", refuseFriend)
	router.POST("/friend/remove", removeFriend)

	// 黑名单
	router.POST("/blacklist/add", addToBlackList)
}

// @Title rcToken
// @Description 获取融云token
// @Accept  json
// @Param   refresh  query   string  true   "是否需要向融云服务器获取新的token（一般填0，注意是字符串）"
// @Success 200 {object} models.S2C_RcToken "融云token"
// @Failure 400 {object} models.APIError "获取融云token失败"
// @Resource /socially
// @Router /socially/rctoken [get]
func getRongCloudToken(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	refresh := c.Query("refresh")

	rcToken, err := apiservice.FindRcToken(userId, refresh)
	if err == nil {
		c.JSON(200, models.S2C_RcToken{rcToken})
	} else {
		c.JSON(400, models.APIError{10411, err.Error()})
	}
}

// @Title list
// @Description 获取好友列表
// @Accept  json
// @Success 200 {object} models.FriendList "返回好友列表"
// @Failure 400 {object} models.APIError "获取好友列表失败"
// @Resource /socially
// @Router /socially/friend/list [get]
func getFriendList(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	friendList, err := apiservice.GetFriendList(userId)
	if err == nil {
		c.JSON(200, friendList)
	} else {
		c.JSON(400, models.APIError{10403, err.Error()})
	}
}

// @Title friendInfo
// @Description 获取好友信息
// @Accept  json
// @Param   targetid  query   string  true   "要查询的用户id"
// @Param   updatetime  query   int64  true   "更新时间"
// @Success 200 {object} models.S2C_UserData "返回好友信息"
// @Failure 400 {object} models.APIError "返回好友信息失败"
// @Resource /socially
// @Router /socially/friend/info [get]
func getFriendInfo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	targetId := c.DefaultQuery("targetid", "")
	updateTime, _ := strconv.ParseInt(c.DefaultQuery("updatetime", "0"), 10, 64)

	friendInfo, err := apiservice.GetFriendInfo(userId, targetId, updateTime)
	if err == nil {
		c.JSON(200, friendInfo)
	} else {
		c.JSON(400, models.APIError{10415, err.Error()})
	}
}

// @Title add
// @Description 请求加好友
// @Accept  json
// @Param   targetid  form   string  true   "要添加的用户id"
// @Success 201 {string} string "发送添加请求成功"
// @Failure 400 {object} models.APIError "发送添加请求失败"
// @Resource /socially
// @Router /socially/friend/add [post]
func addFriendById(c *gin.Context) {
	targetId := c.PostForm("targetid")
	fromId := c.MustGet("userId").(string)

	ecode := apiservice.AddFriendById(fromId, targetId)
	if ecode == 0 {
		c.JSON(201, "addFriendByContacts success")
	} else {
		c.JSON(400, models.APIError{ecode, "addFriendByContacts fail"})
	}
}

// @Title add
// @Description 通过红包号、手机号搜素加好友
// @Accept  json
// @Param   condition  form   string  true   "要添加的手机号或红包号"
// @Success 201 {string} string "发送添加请求成功"
// @Failure 400 {object} models.APIError "发送添加请求失败"
// @Resource /socially
// @Router /socially/friend/addbysearch [post]
func addFriendByAccount(c *gin.Context) {
	targetAccount := c.PostForm("condition")
	userId := c.MustGet("userId").(string)

	ecode := apiservice.AddFriendByAccount(userId, targetAccount)
	if ecode == 0 {
		c.JSON(201, "addFriendBySearch success")
	} else {
		c.JSON(400, models.APIError{ecode, "addFriendBySearch fail"})
	}
}

// @Title agree
// @Description 同意加好友
// @Accept  json
// @Param   userid  form   string  true   "发起请求的用户id"
// @Param   name  form   string  true   "自己的昵称"
// @Param   portrait  form   string  true   "自己的图像"
// @Success 201 {string} string "加好友失败"
// @Failure 400 {object} models.APIError "请求失败"
// @Resource /socially
// @Router /socially/friend/agree [post]
func agreeFriend(c *gin.Context) {
	userIdA := c.PostForm("userid")         // 用户A的id
	userIdB := c.MustGet("userId").(string) // 自己的用户Id

	ecode := apiservice.AgreeFriend(userIdB, userIdA)
	fmt.Println("agree friend ecode = ", ecode)
	if ecode == 0 {
		c.JSON(201, "agreeFriend success")
	} else {
		c.JSON(400, models.APIError{ecode, "agreeFriend fail"})
	}
}

// @Title refuse
// @Description 拒绝添加好友请求
// @Accept  json
// @Param   fromid  form   string  true   "要拒绝的用户id"
// @Success 201 {object} string "成功拒绝"
// @Resource /socially
// @Router /socially/friend/refuse [post]
func refuseFriend(c *gin.Context) {
	userIdA := c.PostForm("fromid")         // A
	userIdB := c.MustGet("userId").(string) // B 被添加的用户

	apiservice.RefuseFriend(userIdB, userIdA)
	c.JSON(201, "refuseFriend success")
}

// @Title remove
// @Description 解除好友关系
// @Accept  json
// @Param   friendid  form   string  true   "好友用户id"
// @Success 201 {string} string "解除成功"
// @Failure 400 {object} models.APIError "删除好友失败"
// @Resource /socially
// @Router /socially/friend/remove [post]
func removeFriend(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	friendId := c.PostForm("friendid")

	if len(friendId) == 0 {
		c.JSON(400, models.APIError{10413, "removeFriend fail"})
	} else {
		apiservice.RemoveFriend(userId, friendId)
		c.JSON(201, "removeFriend success")
	}
}

// @Title addblacklist
// @Description 解除好友关系
// @Accept  json
// @Param   targetid  form   string  true   "用户id"
// @Success 201 {string} string "加入成功"
// @Failure 400 {object} models.APIError "删除好友失败"
// @Resource /socially
// @Router /socially/blacklist/add [post]
func addToBlackList(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	targetid := c.PostForm("targetid")

	if len(targetid) == 0 {
		c.JSON(400, models.APIError{10414, "addToBlackList fail"})
	} else {
		apiservice.AddToBlackList(userId, targetid)
		c.JSON(201, "addToBlackList success")
	}
}
