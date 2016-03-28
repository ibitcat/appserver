package weixin

import (
	"encoding/json"
	"errors"
	"fmt"

	"app-server/models"
	"app-server/pkg/httplib"
)

const (
	AppId     = "wx4f2bab4031677dad"
	AppSecret = "22dcc3e83e78fa035b8084c50c0adb0b"
)

func GetAccessToken(code string) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token"
	req := httplib.Get(url)
	req.Param("appid", AppId)
	req.Param("secret", AppSecret)
	req.Param("code", code)
	req.Param("grant_type", "authorization_code")
	byteData, err := req.Bytes()

	fmt.Println(string(byteData), err)
}

func CheckWeixinAuth(openId, token string) error {
	url := "https://api.weixin.qq.com/sns/auth"
	req := httplib.Get(url)
	req.Param("access_token", token)
	req.Param("openid", openId)
	byteData, err := req.Bytes()

	fmt.Println(string(byteData), err)
	if err != nil {
		return err
	}

	var backMsg struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	jsonErr := json.Unmarshal(byteData, &backMsg)
	if jsonErr != nil {
		return jsonErr
	}

	if backMsg.ErrCode != 0 {
		return errors.New(backMsg.ErrMsg)
	}

	return nil
}

func GetWeixinUserInfo(openId, token string) *models.WeixinUserInfo {
	url := "https://api.weixin.qq.com/sns/userinfo"
	req := httplib.Get(url)
	req.Param("access_token", token)
	req.Param("openid", openId)
	byteData, err := req.Bytes()

	fmt.Println("weixin userinfo--------> ", string(byteData), err)
	if err != nil {
		return nil
	}

	var userInfo models.WeixinUserInfo
	if json.Unmarshal(byteData, &userInfo) != nil {
		return nil
	}

	return &userInfo
}
