// 极光推送

package jpush

import (
	"fmt"

	"app-server/pkg/sdk/jpush/jpushclient"
)

const (
	appKey = "f708414bc10c6c45c6eb4d30"
	secret = "c64221bbec68fb07e87b923f"
)

var jpClient *jpushclient.PushClient

func InitJpushSDK() {
	jpClient = jpushclient.NewPushClient(secret, appKey)
}

// 推送远程通知
func PushNotice() {
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()

	//Audience
	var ad jpushclient.Audience
	s := []string{"56838c10a5129529d0000004"}
	//ad.SetTag(s)
	ad.SetAlias(s)
	//ad.SetID(s)
	//ad.All()

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

	var msg jpushclient.Message
	msg.Title = "Hello"
	msg.Content = "你是ylywn"

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	str, err := jpClient.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}
}
