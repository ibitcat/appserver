// Sub API Definitions, 子模块定义，每个资源定义一次
// 注册user组的路由

package v1

import (
	"fmt"
	"strconv"

	"app-server/apiservice"
	"app-server/define"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 用户操作 [/user]
func Users(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/user")

	// 登陆、注册不需要token验证
	router.POST("/login", login)
	router.POST("/register", register)
	router.GET("/newtoken", refreshTokens)

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/logout", logout)
	router.POST("/nickname", updateNickname)
	router.POST("/account", bindAccount)
	router.POST("/phone", bindPhoneNumber)
	router.POST("/resetphone", resetPhoneNumber)
	router.POST("/sex", updateSex)
	router.POST("/area", updateArea)
	router.POST("/signature", updateSignature)
	router.POST("/portrait", updatePortrait)
	router.GET("/userinfo", getUserInfo)
	router.POST("/resetpwd", resetPassword)
	router.POST("/verifypwd", verifyPassword)

	// 红包相关

	// 排行榜
	router.GET("/rank", getRankList)

	// 系统通知
	router.GET("/sysnotice", getSystemNoticeList)
}

// @Title login
// @Description 用户登录
// @Accept  json
// @Param   username  form   string  true   "手机号码/账号"
// @Param   password  form   string  true   "用户密码"
// @Success 201 {object} models.S2C_TokenArray "用户登录成功"
// @Failure 400 {object} models.APIError "用户登录失败"
// @Resource /user
// @Router /user/login [post]
func login(c *gin.Context) {
	account := c.PostForm("username")
	password := c.PostForm("password")

	ecode, accessToken, refreshToken := apiservice.Login(account, password)
	if ecode == 0 {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{ecode, "login fail"})
	}
}

// @Title logout
// @Description 用户退出
// @Accept  json
// @Success 201 {string} string "用户退出成功"
// @Failure 400 {object} models.APIError "退出错误"
// @Resource /user
// @Router /user/logout [post]
func logout(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	err := apiservice.Logout(c.Request, userId)
	fmt.Println("logout -->", err)
	if err == nil {
		c.JSON(201, "logout success")
	} else {
		c.JSON(400, models.APIError{10005, err.Error()})
	}
}

// @Title register
// @Description 用户手机号码注册
// @Accept  json
// @Param   phonenum  form   string  true   "手机号码"
// @Param   password  form   string  true   "用户密码"
// @Success 201 {object} models.S2C_TokenArray "用户登录成功"
// @Failure 400 {object} models.APIError "用户注册失败"
// @Resource /user
// @Router /user/register [post]
func register(c *gin.Context) {
	phonenum := c.PostForm("phonenum")
	password := c.PostForm("password")

	ecode, accessToken, refreshToken := apiservice.Register(phonenum, password)
	if ecode == 0 {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{ecode, "register fail"})
	}
}

// @Title userinfo
// @Description 拉取用户的基本信息
// @Accept  json
// @Param   updatetime  query   int64  true   "更新时间"
// @Success 200 {object} models.S2C_UserData "拉取用户数据成功"
// @Failure 400 {object} models.APIError "更新失败"
// @Resource /user
// @Router /user/userinfo [get]
func getUserInfo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	updateTime, _ := strconv.ParseInt(c.DefaultQuery("updatetime", "0"), 10, 64)

	user, err := apiservice.GetUserInfo(userId, updateTime)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(400, models.APIError{10010, err.Error()})
	}
}

// @Title resetpassword
// @Description 重置密码
// @Accept  json
// @Param   password  form   string  true   "新的密码"
// @Param   flag  form   string  true   "操作类型 1=找回密码 0=修改密码"
// @Param   code  form   string  true   "验证码"
// @Success 201 {object} models.S2C_TokenArray "重置密码成功"
// @Failure 400 {object} models.APIError "重置密码失败"
// @Resource /user
// @Router /user/resetpwd [post]
func resetPassword(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	password := c.PostForm("password")
	flag := c.PostForm("flag")
	code := c.PostForm("code")

	ecode, accessToken, refreshToken := apiservice.ResetPassword(userId, password, flag, code)
	if ecode == 0 {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{ecode, "reset password fail"})
	}
}

// @Title verifypassword
// @Description 验证密码
// @Accept  json
// @Param   password  form   string  true   "原密码"
// @Success 201 {string} string "验证密码成功"
// @Failure 400 {object} models.APIError "验证密码失败"
// @Resource /user
// @Router /user/verifypwd [post]
func verifyPassword(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	password := c.PostForm("password")

	ecode := apiservice.VerifyPassword(userId, password)
	if ecode == 0 {
		c.JSON(201, "verify password success")
	} else {
		c.JSON(400, models.APIError{ecode, "verify password fail"})
	}
}

// @Title newtoken
// @Description 刷新token，用refresh token来刷新
// @Accept  json
// @Param   refresh_token  query   string  true   "刷新token"
// @Success 200 {object} models.S2C_TokenArray "换取新的token成功"
// @Failure 400 {object} models.APIError "换取token失败"
// @Resource /user
// @Router /user/newtoken [get]
func refreshTokens(c *gin.Context) {
	tokenStr := c.Query("refresh_token")

	ecode, accessToken, refreshToken := apiservice.GenerateNewToken(tokenStr)
	if ecode == 0 {
		c.JSON(200, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{ecode, "refreshTokens fail"})
	}
}

// @Title account
// @Description 设置红包号，只能设置一次
// @Accept  json
// @Param   account  form   string  true   "红包账号名"
// @Success 201 {string} string "设置红包账号成功"
// @Failure 400 {object} models.APIError "设置红包账号失败"
// @Resource /user
// @Router /user/account [post]
func bindAccount(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	account := c.PostForm("account")

	ecode := apiservice.BindAccount(userId, account)
	if ecode == 0 {
		c.JSON(201, "bind account success")
	} else {
		c.JSON(400, models.APIError{ecode, "bind account fail"})
	}
}

// @Title bindphone
// @Description 绑定手机号并设置密码
// @Accept  json
// @Param   phonenum  form   string  true   "手机号"
// @Param   password  form   string  true   "密码"
// @Success 201 {string} string "绑定成功"
// @Failure 400 {object} models.APIError "绑定失败"
// @Resource /user
// @Router /user/phone [post]
func bindPhoneNumber(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	phonenum := c.PostForm("phonenum")
	password := c.PostForm("password")

	ecode := apiservice.BindPhoneNum(userId, phonenum, password)
	if ecode == 0 {
		c.JSON(201, "bindPhoneNumber success")
	} else {
		c.JSON(400, models.APIError{ecode, "bindPhoneNumber fail"})
	}
}

// @Title resetphone
// @Description 更换手机号码
// @Accept  json
// @Param   phonenum  form   string  true   "手机号"
// @Param   password  form   string  true   "密码"
// @Param   code  form   string  true   "手机验证码"
// @Success 201 {string} string "更换成功"
// @Failure 400 {object} models.APIError "跟换失败"
// @Resource /user
// @Router /user/resetphone [post]
func resetPhoneNumber(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	phonenum := c.PostForm("phonenum")
	password := c.PostForm("password")
	code := c.PostForm("code")

	ecode := apiservice.ResetPhoneNum(userId, phonenum, password, code)
	if ecode == 0 {
		c.JSON(201, "resetPhoneNumber success")
	} else {
		c.JSON(400, models.APIError{ecode, "resetPhoneNumber fail"})
	}
}

// @Title nickname
// @Description 更改用户昵称
// @Accept  json
// @Param   nickname  form   string  true   "用户昵称"
// @Success 201 {string} string "修改昵称成功"
// @Failure 400 {object} models.APIError "修改昵称失败"
// @Resource /user
// @Router /user/nickname [post]
func updateNickname(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	nickName := c.PostForm("nickname")

	err := apiservice.UpdatePersonalInfo(userId, define.EUser_Nickname, nickName)
	if err == nil {
		c.JSON(201, "update success")
	} else {
		c.JSON(400, models.APIError{10004, "update fail"})
	}
}

// @Title area
// @Description 设置性别
// @Accept  json
// @Param   sex  form   int  true   "性别1=男0=女"
// @Success 201 {string} string "设置性别成功"
// @Resource /user
// @Router /user/sex [post]
func updateSex(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	sex, _ := strconv.Atoi(c.DefaultPostForm("sex", "0"))

	apiservice.UpdatePersonalInfo(userId, define.EUser_Sex, sex)
	c.JSON(201, "update sex success")
}

// @Title area
// @Description 设置区域
// @Accept  json
// @Param   area  body   models.AreaInfo  true   "所在区域"
// @Success 201 {objson} string "设置区域成功"
// @Failure 400 {object} models.APIError "设置区域失败"
// @Resource /user
// @Router /user/area [post]
func updateArea(c *gin.Context) {
	var area models.AreaInfo
	userId := c.MustGet("userId").(string)
	err := c.BindJSON(&area)
	//b, _ := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		err = apiservice.UpdatePersonalInfo(userId, define.EUser_Area, area)
	}

	if err == nil {
		c.JSON(201, "update area success")
	} else {
		c.JSON(400, models.APIError{10015, err.Error()})
	}
}

// @Title signature
// @Description 设置个性签名
// @Accept  json
// @Param   signature  form   string  true   "个性签名"
// @Success 201 {string} string "设置个性签名成功"
// @Failure 400 {object} models.APIError "设置个性签名失败"
// @Resource /user
// @Router /user/signature [post]
func updateSignature(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	signature := c.PostForm("signature")

	err := apiservice.UpdatePersonalInfo(userId, define.EUser_Signature, signature)
	if err == nil {
		c.JSON(201, "update signature success")
	} else {
		c.JSON(400, models.APIError{10015, err.Error()})
	}
}

// @Title portrait
// @Description 设置个人头像
// @Accept  json
// @Param   portrait_key  form   string  true   "七牛上的文件key"
// @Success 201 {string} string "设置头像成功"
// @Failure 400 {object} models.APIError "设置头像失败"
// @Resource /user
// @Router /user/portrait [post]
func updatePortrait(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	portrait := c.PostForm("portrait_key")

	err := apiservice.UpdatePersonalInfo(userId, define.EUser_Portrait, portrait)
	if err == nil {
		c.JSON(201, "update signature success")
	} else {
		c.JSON(400, models.APIError{10015, err.Error()})
	}
}

// @Title rankList
// @Description 排行榜
// @Accept  json
// @Param   ranktype  query   int  true   "排行榜类型[1=红包榜 2=好友榜 3=等级榜 4=老板榜]"
// @Param   page  query   int  true   "页数[0开始]"
// @Success 200 {object} models.S2C_RankList "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /user
// @Router /user/rank [get]
func getRankList(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	page := c.Query("page")
	ranktype := c.Query("ranktype")

	ret := apiservice.GetRankListByType(userId, ranktype, page)
	if ret != nil {
		c.JSON(200, ret)
	} else {
		c.JSON(400, models.APIError{10004, "getRankList fail"})
	}
}

// @Title noticeList
// @Description 系统通知列表
// @Accept  json
// @Param   page  query   int  true   "页数[0开始]"
// @Success 200 {object} models.S2C_SysNoticeList "成功"
// @Failure 400 {object} models.APIError "失败"
// @Resource /user
// @Router /user/sysnotice [get]
func getSystemNoticeList(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	ret := apiservice.GetSystemNoticeList(userId)
	if ret != nil {
		c.JSON(200, ret)
	} else {
		c.JSON(400, models.APIError{10004, "getSystemNoticeList fail"})
	}
}
