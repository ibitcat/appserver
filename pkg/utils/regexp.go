/*
 正则匹配相关
*/

package utils

import (
	"regexp"
)

const (
	E_UnallowAble   = iota // 无效账号
	E_AccountPhone         // 手机账号
	E_AccountString        // 红包账号
)

// 是否是手机号码,暂只支持国内手机号码 "13877778888"
func IsPhoneNumber(phoneNum string) bool {
	reg := regexp.MustCompile(`^1([38][0-9]|14[57]|5[^4])\d{8}$`)
	return reg.MatchString(phoneNum)
}

// 是否是正确的账号（英文开头并且只能包含英文字符和数字）
func IsAccount(account string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)
	return reg.MatchString(account)
}

// 检查用户名
func CheckAccout(account string) uint8 {
	if IsPhoneNumber(account) {
		return E_AccountPhone
	}

	if IsAccount(account) {
		return E_AccountString
	}

	return E_UnallowAble
}

// 是否为合法的账号
func IsAllowableAccout(account string) bool {
	accountType := CheckAccout(account)
	return accountType > E_UnallowAble && accountType <= E_AccountString
}
