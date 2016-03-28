// 用户等级配置

package config

import (
	"bufio"
	"encoding/json"
	"os"
)

type UserLvItem struct {
	Level       int `json:"lv"`            // 等级
	TotalPoint  int `json:"totalintegral"` // 升到该等级需要的总经验
	RedpktLimit int `json:"redenvelope"`   // 每天抢红包的次数
}

var (
	userlvcfg_path string = "res/json/LevelTable.json"
	UserLevelCfg   []UserLvItem
)

// 解析
func parseUserLevelConfig() {
	cfgFile, err := os.Open(userlvcfg_path)
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

	var temp map[string]UserLvItem
	err = json.Unmarshal(bytes, &temp)
	if err != nil {
		panic(err)
	}

	l := len(temp)
	if l > 0 {
		UserLevelCfg = make([]UserLvItem, l, l)
		for _, v := range temp {
			UserLevelCfg[v.Level] = v
		}
	}
}

func GetUserLevelByPoint(point int) int {
	var level int
	for _, v := range UserLevelCfg {
		if point > v.TotalPoint {
			level++
		} else {
			break
		}
	}
	if level >= len(UserLevelCfg) {
		level = len(UserLevelCfg) - 1
	}
	return level
}

func GetGrabLimit(lv int) int {
	if lv < 0 {
		return 0
	}

	if lv >= len(UserLevelCfg) {
		lv = len(UserLevelCfg) - 1
	}

	return UserLevelCfg[lv].RedpktLimit
}
