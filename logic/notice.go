// 通知

package logic

import (
	"encoding/json"
	"fmt"
	"time"

	"app-server/define"
	"app-server/models"
	"app-server/pkg/redis"
)

// push系统通知
func PushSystemNotice(userId, txt string) {
	if len(userId) > 0 {
		notice := models.SysNoticeItem{}
		notice.Text = txt
		notice.Time = time.Now().Unix()
		notice.Flag = 0

		bytes, err := json.Marshal(notice)
		if err != nil {
			return
		}

		key := define.SystemNotice + userId
		redis.Do("RPUSH", key, string(bytes))
	}
}

// push远程推送
func PushRemoteNotice(userId, txt string) {
	notice := models.RmtNoticeItem{}
	notice.UId = userId
	notice.Time = time.Now().Unix()
	notice.Text = txt

	bytes, err := json.Marshal(notice)
	if err != nil {
		return
	}
	redis.Do("RPUSH", define.RemoteNotice, string(bytes))
}

// 获取用户的系统通知
func GetSystemNoticeList(userId string) []models.SysNoticeItem {
	list := make([]models.SysNoticeItem, 0, 10)
	key := define.SystemNotice + userId
	values, err := redis.GetStrings("LRANGE", key, 0, -1)
	for _, v := range values {
		var notice models.SysNoticeItem
		err = json.Unmarshal([]byte(v), &notice)
		if err != nil {
			continue
		}

		list = append(list, notice)
	}

	return list
}

// 获取红包的系统通知
func GetRedpacketNotice(userId string) []models.SysNoticeItem {
	now := time.Now().Unix()
	idList, err := redis.GetStrings("ZRANGEBYSCORE", define.RedpktListKey, 0, now)
	if err != nil {
		fmt.Println("获取红包通知列表 err =  ", err)
		return nil
	}

	list := make([]models.SysNoticeItem, 0, 10)
	for _, id := range idList {
		con, ok := RedpacketContainers[id]
		if !ok {
			continue
		}

		if !con.IsStart() {
			continue
		}

		switch c := con.(type) {
		case *ShareRedpacket: // 分享类红包
			if gs := c.getGrabStatus(userId); gs != nil {
				if now >= gs.Expire { // 过期
					notice := models.SysNoticeItem{}
					if gs.Status == 1 {
						notice.Text = "在规定时间内未分享到朋友圈，任务失败您失去XXXX的红包！"
					} else if gs.Status == 2 {
						notice.Text = "在规定时间内未上传截图确认成功，任务失败您失去XXXX的红包，你的待确认资金减少X.XX元！"
					}
					notice.Time = gs.Expire
					notice.Flag = 0
					list = append(list, notice)
				}
			}
		case *AppRedpacket: // app导量类红包
			if gs := c.getGrabStatus(userId); gs != nil {
				if now >= gs.Expire { // 过期
					notice := models.SysNoticeItem{}
					notice.Text = "在规定时间内未下载并试玩游戏，任务失败您失去XXXX的红包！if "
					notice.Time = gs.Expire
					notice.Flag = 0
					list = append(list, notice)
				}
			}
		default:
			continue
		}
	}

	return list
}
