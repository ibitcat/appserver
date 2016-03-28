// 消息和推送

package models

// 系统通知节点
type SysNoticeItem struct {
	Text string `json:"text"` // 通知文本
	Time int64  `json:"time"` // 通知时间
	Flag int    `json:"flag"` // 是否已读
}

// 远程推送节点
type RmtNoticeItem struct {
	UId  string `json:"uid"`  // 推送目标id
	Text string `json:"text"` // 推送文本
	Time int64  `json:"time"` // 推送时间
}

/////////////////////////////////////////////////////////
type S2C_SysNoticeList struct {
	List []SysNoticeItem `json:"list,omitempty"`
}
