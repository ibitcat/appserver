/*
	redis key 定义
*/

package define

const (
	AddFriendPrefix    = "addFriend:"       // 添加好友缓存key的前缀
	UserCachePrefix    = "userCache:"       // 用户数据缓存key的前缀
	RedpktListKey      = "redpacketlist"    // 红包列表key
	RedpktGrabRrefix   = "redpacketgrab:"   // 正在抢红包的记录key前缀
	RedpktRecordPrefix = "redpacketrecord:" // 已抢到红包的记录key前缀
	RedpktDevicePrefix = "redpacketdevice:" // 抢红包用户和设备id的映射
	GrabStatusPrefix   = "grabstatus:"      // 用户抢红包的状态
	BeenSendPrefix     = "beensend:"        // 用户已发送的红包列表
	IncomeRank         = "income_rank"      // 用户总收入排名（红包排行榜）
	OutcomeRank        = "outcome_rank"     // 用户总收入排名（老板排行榜）
	LevelRank          = "level_rank"       // 用户等级排行榜
	FriendRank         = "friend_rank"      // 好友排行榜
	SystemNotice       = "system_notice:"   // 系统通知
	RemoteNotice       = "remote_notice"    // 远程推送(调用极光推送api)
)
