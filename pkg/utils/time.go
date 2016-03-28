package utils

import (
	"math"
	"strconv"
	"time"
)

// 计算两个日期的相隔天数
// 时间格式 "20151021",表示2015年10月21日
func IntervalDays(startDate, stopDate uint32) uint32 {
	var startDay, stopDay time.Time
	var err error
	startDay, err = time.ParseInLocation("20060102", strconv.FormatUint(uint64(startDate), 10), time.Local)
	if err != nil {
		return 0
	}

	stopDay, err = time.ParseInLocation("20060102", strconv.FormatUint(uint64(stopDate), 10), time.Local)
	if err != nil {
		return 0
	}
	days := uint32(math.Abs(stopDay.Sub(startDay).Seconds()) / 86400)
	return days + 1
}
