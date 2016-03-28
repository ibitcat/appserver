// token model

package models

/////////////////////////////////////////////////////////
// appserver 的token
type TokenInfo struct {
	Token     string `description:"token字符串"`
	Expiresin uint32 `description:"token的有效期，以s为单位"`
}

/////////////////////////////////////////////////////////
// 返回给客户端
// token数组
type S2C_TokenArray struct {
	AccessToken  TokenInfo `json:"access_token" description:"access token信息"`
	RefreshToken TokenInfo `json:"refresh_token" description:"refresh token信息"`
}

// 返回给客户端的融云token
type S2C_RcToken struct {
	RcToken string `josn:"rctoken"`
}

/////////////////////////////////////////////////////////
// 七牛的上传凭证
type S2C_QiniuUpToken struct {
	UploadToken string `json:"upload_token"`
	Expires     int    `json:"expires"`
}

// 七牛私有空间下载链接
type S2C_QiniuDlUrl struct {
	DownloadUrl string `json:"dl_url"`
	Expires     int    `json:"expires"`
}
