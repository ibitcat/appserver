// 红包相关配置
package config

var RedpktLimitCfg map[int]int = map[int]int{1: 2, 4: 4}           // 每种红包每天可以完成的次数
var RedpktOwnCfg map[int]int = map[int]int{1: 1, 2: 1, 3: 1, 4: 2} // 每种红包最多可拥有的次数

// 根据红包类型获取每日领取次数限制
func GetLimitByType(tp int) int {
	limit, ok := RedpktLimitCfg[tp]
	if ok {
		return limit
	}
	return 0
}

// 根据红包类型获取每种红包最多拥有的次数限制
func GetOwnLimitByType(tp int) int {
	own, ok := RedpktOwnCfg[tp]
	if ok {
		return own
	}
	return 0
}
