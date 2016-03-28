// 刷新token的操作定义
package define

const (
	ERefreshToken_Refresh  = iota + 1 // 用refresh token 刷新
	ERefreshToken_Login               // 登陆刷新token
	ERefreshToken_Forcibly            // 强制刷新access 和 refresh token，比如修改密码
)
