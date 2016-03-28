// 用户积分配置

package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type UserPointItem struct {
	Id    int `json:"bh"`
	Limit int `json:"dailylimit"`     // 次数限制 0=无次数限制
	Point int `json:"singleintegral"` // 积分
}

var (
	userpointcfg_path string = "res/json/UserPoint.json"
	UserPointCfg      map[int]UserPointItem
)

// 读取用户积分配置
func parseUserPointConfig() {
	cfgFile, err := os.Open(userpointcfg_path)
	if err != nil {
		panic(err)
	}
	defer cfgFile.Close()

	fi, _ := cfgFile.Stat()
	var size int64 = fi.Size()
	bytes := make([]byte, size)

	// read
	buf := bufio.NewReader(cfgFile)
	_, err = buf.Read(bytes)
	if err != nil {
		panic(err)
	}

	var temp map[string]UserPointItem
	err = json.Unmarshal(bytes, &temp)
	if err != nil {
		panic(err)
	}

	l := len(temp)
	if l > 0 {
		UserPointCfg = make(map[int]UserPointItem)
		for _, v := range temp {
			UserPointCfg[v.Id] = v
		}
	}

	fmt.Println("用户积分配置 = ", UserPointCfg)
}

// 获取积分配置
func GetPointCfg(tp int) *UserPointItem {
	item, ok := UserPointCfg[tp]
	if ok {
		return &item
	}

	return nil
}
