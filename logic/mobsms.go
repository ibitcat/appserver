package logic

import (
	"crypto/tls"
	"errors"
	"fmt"

	"app-server/pkg/httplib"
	"app-server/pkg/utils"
)

// 验证短信验证码
func VerifySmsCode(phone, zone, code string) error {
	// mob短信验证码
	mobApiUrl := "https://web.sms.mob.com/sms/verify"
	req := httplib.Post(mobApiUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //跳过证书验证

	appkey := "d6bd91387dcc"
	if len(appkey) > 0 && len(phone) > 0 && len(zone) > 0 && len(code) > 0 {
		req.Param("appkey", appkey)
		req.Param("phone", phone)
		req.Param("zone", zone)
		req.Param("code", code)

		retcode, err := req.String()
		if err != nil {
			return err
		}

		json, jsonErr := utils.JsonDecode(retcode)
		if jsonErr != nil {
			return jsonErr
		}

		ecode := json.(map[string]interface{})["status"]
		if value, ok := ecode.(float64); ok {
			fmt.Println("mob sms code = ", value)
			if uint32(value) == 200 {
				return nil
			}
		}
	}

	return errors.New("verify sms code fail")
}
