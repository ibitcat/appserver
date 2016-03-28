/*
	数学相关
*/

package utils

import (
	"math"
	"math/rand"
	"time"
)

// 四舍五入
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

// 区间随机数 [min,max)
func RandomInterval(min, max int) int {
	// 用当前时间戳作为随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min == max || min > max {
		return min
	}

	return min + r.Int()%(max-min)
}
