package qiniu

import (
	"fmt"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

var (
	qiniuBucket kodo.Bucket  // 七牛工作空间
	qiniuClient *kodo.Client // 七牛client对象
)

const (
	bucketName          = "domi"                                     // 工作空间名称
	privateBucketName   = "private"                                  // 私人工作空间名称
	bucketDomain        = "7xo7xm.com2.z0.glb.qiniucdn.com"          // 工作空间域名
	privateBucketDomain = "7xpgcu.com2.z0.glb.qiniucdn.com"          // 私人工作空间域名
	accessKey           = "jGYeZ-D_hmaqUqxmD7M3O6_WoLlrm1a2Pi5kn-xY" // access key
	secretKey           = "8r23L_EvXYF-ckrbJiPvWivbiWS2AjQSs8pYFR8O" // secret key （需防止泄露）
)

// 初始化七牛SDK
func InitQiniuSdk() bool {
	kodo.SetMac(accessKey, secretKey)
	qiniuClient = kodo.New(0, nil) // 创建client
	if qiniuClient == nil {
		return false
	}

	qiniuBucket = qiniuClient.Bucket(bucketName)
	return true
}

func MakeUploadToken(key string) string {
	fmt.Println("生成qiniu 上传凭证------->")
	policy := new(kodo.PutPolicy)
	policy.Scope = bucketName + ":" + key
	policy.Expires = 3600
	// policy.CallbackUrl = "http://domiapp.ngrok.natapp.cn/v1/pub/posttest"
	// policy.CallbackHost = "domiapp.ngrok.natapp.cn"
	// policy.CallbackBody = "name=$(fname)&hash=$(etag)"
	// policy.CallbackFetchKey = 1

	return qiniuClient.MakeUptoken(policy)
}

func MakePrivateUploadToken(key string) string {
	fmt.Println("生成私人qiniu 上传凭证------->")
	policy := new(kodo.PutPolicy)
	policy.Scope = privateBucketName + ":" + key
	policy.Expires = 3600
	// policy.CallbackUrl = "http://domiapp.ngrok.natapp.cn/v1/pub/posttest"
	// policy.CallbackHost = "domiapp.ngrok.natapp.cn"
	// policy.CallbackBody = "name=$(fname)&hash=$(etag)"
	// policy.CallbackFetchKey = 1

	return qiniuClient.MakeUptoken(policy)
}

func MakePrivateUrl(key string) string {
	baseUrl := kodo.MakeBaseUrl(privateBucketDomain, key)
	return qiniuClient.MakePrivateUrl(baseUrl, nil)
}

// 七牛上传测试
func TestUpload(tokenStr string) {
	ctx := context.Background()

	err := qiniuBucket.PutFileWithToken(ctx, nil, "1.jpg", tokenStr, nil)
	if err != nil {
		fmt.Println("上传文件失败,err = ", err)
		return
	}
}
