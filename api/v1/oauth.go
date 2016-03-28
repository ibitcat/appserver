// 第三方登陆

package v1

import (
	"app-server/apiservice"
	"app-server/models"
	"app-server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @SubApi 第三方认证 [/oauth]
func Oauth(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/oauth")

	router.POST("/weixin/login", loginByWeixin)
	router.POST("/qq/login", loginByQQ)
	router.POST("/weibo/login", loginByWeibo)

	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/bindweixin", loginByWeixin)
	router.POST("/bindqq", loginByQQ)
	router.POST("/bindweibo", loginByWeibo)
	router.POST("/unbind", unbindOauthAccount)
}

// @Title loginByweixin
// @Description 使用微信登陆
// @Accept  json
// @Param   openid  form   string  true   "微信openid"
// @Param   token  form   string  true   "微信access_token"
// @Success 201 {object} models.S2C_TokenArray "登陆成功"
// @Failure 400 {object} models.APIError "登陆失败"
// @Resource /oauth
// @Router /oauth/weixin/login [post]
func loginByWeixin(c *gin.Context) {
	openId := c.PostForm("openid")
	tokenStr := c.PostForm("token")

	err, accessToken, refreshToken := apiservice.LoginByWeixin(openId, tokenStr)
	if err == nil {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{10016, err.Error()})
	}
}

// @Title loginbyQQ
// @Description 使用QQ登陆
// @Accept  json
// @Param   openid  form   string  true   "qq openid"
// @Param   openkey  form   string  true   "qq openkey"
// @Success 201 {object} models.S2C_TokenArray "登陆成功"
// @Failure 400 {object} models.APIError "登陆失败"
// @Resource /oauth
// @Router /oauth/qq/login [post]
func loginByQQ(c *gin.Context) {
	openId := c.PostForm("openid")
	openKey := c.PostForm("openkey")

	err, accessToken, refreshToken := apiservice.LoginByQQ(openId, openKey)
	if err == nil {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{10017, err.Error()})
	}
}

// @Title loginbyWebo
// @Description 使用微博登陆
// @Accept  json
// @Param   access_token  form   string  true   "weibo的access token"
// @Success 201 {object} models.S2C_TokenArray "登陆成功"
// @Failure 400 {object} models.APIError "登陆失败"
// @Resource /oauth
// @Router /oauth/weibo/login [post]
func loginByWeibo(c *gin.Context) {
	weiboToken := c.PostForm("access_token")

	err, accessToken, refreshToken := apiservice.LoginByWeibo(weiboToken)
	if err == nil {
		c.JSON(201, models.S2C_TokenArray{*accessToken, *refreshToken})
	} else {
		c.JSON(400, models.APIError{10018, err.Error()})
	}
}

// @Title bindweixin
// @Description 绑定微信账号
// @Accept  json
// @Param   openid  form   string  true   "微信openid"
// @Param   token  form   string  true   "微信access_token"
// @Success 201 {string} string "绑定成功"
// @Failure 400 {object} models.APIError "绑定失败"
// @Resource /oauth
// @Router /oauth/bindweixin [post]
func bindWeixin(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	openId := c.PostForm("openid")
	tokenStr := c.PostForm("token")

	ecode := apiservice.BindWeixinAccount(userId, openId, tokenStr)
	if ecode == 0 {
		c.JSON(201, "bindWeixin ok")
	} else {
		c.JSON(400, models.APIError{ecode, "bindWeixin fail"})
	}
}

// @Title bindqq
// @Description 绑定QQ账号
// @Accept  json
// @Param   openid  form   string  true   "qq openid"
// @Param   openkey  form   string  true   "qq openkey"
// @Success 201 {string} string "绑定成功"
// @Failure 400 {object} models.APIError "绑定失败"
// @Resource /oauth
// @Router /oauth/bindqq [post]
func bindQQ(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	openId := c.PostForm("openid")
	openKey := c.PostForm("openkey")

	ecode := apiservice.BindQQAccount(userId, openId, openKey)
	if ecode == 0 {
		c.JSON(201, "bindQQ ok")
	} else {
		c.JSON(400, models.APIError{ecode, "bindQQ fail"})
	}
}

// @Title bindWebo
// @Description 绑定微博账号
// @Accept  json
// @Param   access_token  form   string  true   "weibo的access token"
// @Success 201 {string} string "绑定成功"
// @Failure 400 {object} models.APIError "绑定失败"
// @Resource /oauth
// @Router /oauth/bindweibo [post]
func bindWeibo(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	weiboToken := c.PostForm("access_token")

	ecode := apiservice.BindWeiboAccount(userId, weiboToken)
	if ecode == 0 {
		c.JSON(201, "bindWeibo ok")
	} else {
		c.JSON(400, models.APIError{ecode, "bindWeibo fail"})
	}
}

// @Title unbind
// @Description 解绑第三方账号
// @Accept  json
// @Param   platform  form   string  true   "第三方平台[0=微信,1=微博,2=QQ]"
// @Success 201 {string} string "解绑成功"
// @Failure 400 {object} models.APIError "绑定失败"
// @Resource /oauth
// @Router /oauth/unbind [post]
func unbindOauthAccount(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	platform := c.PostForm("platform")

	ecode := apiservice.UnbindOauthAccount(userId, platform)
	if ecode == 0 {
		c.JSON(201, "unbind ok")
	} else {
		c.JSON(400, models.APIError{ecode, "unbindOauthAccount fail"})
	}
}
