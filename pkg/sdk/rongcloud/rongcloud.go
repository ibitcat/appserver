/*
融云SDK
*/

package rongcloud

import (
	"log"
)

var RcServer *RCServer // 融云sdk

// 初始化融云SDK
func InitRcSDK() bool {
	var err error
	RcServer, err = NewRCServer("k51hidwq1jqtb", "FAc2Ku1Jq1", "json")
	if err != nil {
		log.Println("初始化融云sdk失败……")
		return false
	}

	return true
}
