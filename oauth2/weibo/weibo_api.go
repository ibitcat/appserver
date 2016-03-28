package weibo

import (
	"encoding/json"
	"fmt"

	"app-server/models"
	"app-server/pkg/httplib"
)

type weiboTokenInfo struct {
	UID      int64  `json:"uid"`
	Appkey   string `json:"appkey"`
	CreateAt int    `json:"create_at"`
	ExpireIn int    `json:"expire_in"`
	//Scope    string `json:"scope"`
}

func CheckWeiboTokenInfo(token string) int64 {
	url := "https://api.weibo.com/oauth2/get_token_info"

	req := httplib.Post(url)
	req.Param("access_token", token)

	byteData, err := req.Bytes()
	fmt.Println("get weibo token info  -------> ", string(byteData), err)
	if err != nil {
		return 0
	}

	var tokenInfo weiboTokenInfo
	jsonErr := json.Unmarshal(byteData, &tokenInfo)
	if jsonErr != nil {
		return 0
	}

	if tokenInfo.UID == 0 {
		return 0
	}

	if tokenInfo.ExpireIn == 0 {
		return 0
	}

	return tokenInfo.UID
}

// 获取微博用户信息
func GetWeiboUserInfo(uidStr, token string) *models.WeiboUserInfo {
	url := "https://api.weibo.com/2/users/show.json"

	req := httplib.Get(url)
	req.Param("access_token", token)
	req.Param("uid", uidStr)

	fmt.Println(req)
	byteData, err := req.Bytes()
	if err != nil {
		return nil
	}

	var userInfo models.WeiboUserInfo
	jsonErr := json.Unmarshal(byteData, &userInfo)
	if jsonErr != nil {
		return nil
	}

	return &userInfo
}
