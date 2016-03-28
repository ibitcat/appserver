// 积分类型,需要与配置保持一致

package define

const (
	EPoint_Login     = iota + 1 // 登陆5次
	EPoint_Redpacket            // 完成红包任务 分享类需要分享成功 导量类需要下载成功
	EPoint_Invite               // 邀请好友
)
