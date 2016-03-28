package qq

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"app-server/models"
	"app-server/pkg/httplib"
)

const (
	AppId  = "1104900236"
	AppKey = "CsdMU8GxB6Za7F1F"
)

// 获取qq api签名
func getQQSignature(urlStr, openId, openKey string) string {
	first := url.QueryEscape(urlStr)

	query := fmt.Sprintf("appid=%s&format=json&openid=%s&openkey=%s&pf=qzone", AppId, openId, openKey)
	second := url.QueryEscape(query)

	signingString := "GET&" + first + "&" + second
	keyBytes := AppKey + "&"

	hasher := hmac.New(sha1.New, []byte(keyBytes))
	hasher.Write([]byte(signingString))

	seg := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(seg)
}

func IsLogin(openId, openKey string) error {
	url := "http://113.108.20.23/v3/user/is_login"
	sig := getQQSignature("/v3/user/is_login", openId, openKey)

	req := httplib.Get(url)
	req.Param("appid", AppId)
	req.Param("sig", sig)
	req.Param("openid", openId)
	req.Param("openkey", openKey)
	req.Param("pf", "qzone")
	req.Param("format", "json")

	byteData, err := req.Bytes()
	fmt.Println("qq is login -------> ", string(byteData), err)
	if err != nil {
		return err
	}

	var backMsg struct {
		RetCode int    `json:"ret"`
		RetMsg  string `json:"msg"`
	}

	jsonErr := json.Unmarshal(byteData, &backMsg)
	if jsonErr != nil {
		return jsonErr
	}

	if backMsg.RetCode != 0 {
		return errors.New(backMsg.RetMsg)
	}

	return nil
}

func GetQQUserInfo(openId, openKey string) *models.QQUserInfo {
	url := "http://113.108.20.23/v3/user/get_info"
	sig := getQQSignature("/v3/user/get_info", openId, openKey)

	req := httplib.Get(url)
	req.Param("appid", AppId)
	req.Param("sig", sig)
	req.Param("openid", openId)
	req.Param("openkey", openKey)
	req.Param("pf", "qzone")
	req.Param("format", "json")

	byteData, err := req.Bytes()
	fmt.Println("QQ userinfo------> ", string(byteData), err)
	if err != nil {
		return nil
	}

	var userInfo models.QQUserInfo
	if json.Unmarshal(byteData, &userInfo) != nil {
		return nil
	}

	return &userInfo
}
