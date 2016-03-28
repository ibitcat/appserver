package apiservice

import (
	"app-server/pkg/sdk/qiniu"
)

// 生成上传凭证，有效期1小时
func MakeQiniuUploadToken(key string) string {
	return qiniu.MakeUploadToken(key)
}

// 生成私有上传凭证，有效期1小时
func MakeQiniuPrivateUploadToken(key string) string {
	return qiniu.MakePrivateUploadToken(key)
}

// 生成私有的下载url
func MakeQiniuPrivateUrl(key string) string {
	if len(key) == 0 {
		return ""
	}
	return qiniu.MakePrivateUrl(key)
}

// 七牛上传测试
func TestQiniuUpload(tokenStr string) {
	qiniu.TestUpload(tokenStr)
}
