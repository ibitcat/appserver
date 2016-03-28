// 排行榜相关

package models

// 排行榜信息
type RankItem struct {
	Rank  int    `json:"rank"`  // 名次
	Score int    `json:"score"` // 分值
	Id    string `json:"id"`    // 用户id
	Name  string `json:"name"`  // 用户昵称
}

// 发送给客户端的排行榜列表
type S2C_RankList struct {
	Type int        `json:"type"` // 排行榜类型 1=红包榜 2=好友榜 3=等级榜 4=老板榜
	List []RankItem `json:"list"` // 排行榜列表
	Self int        `json:"self"` // 自己的排名 0=未上榜，统一表示99999+
}
