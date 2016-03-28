package timer

import (
	"fmt"
	//"sync/atomic"
	"testing"
	"time"
)

var sum int32 = 0
var N int32 = 10
var tt *Timer

func now() {
	fmt.Println("ontimer = ", time.Now().Format("2006-01-02 15:04:05"))
	// atomic.AddInt32(&sum, 1)
	// v := atomic.LoadInt32(&sum)
	// if v == 2*N {
	// 	tt.Stop()
	// }
}

func TestTimer(t *testing.T) {
	timer := NewTimer(time.Millisecond * 10)
	tt = timer
	fmt.Println(timer)
	//var i int32
	//timer.AddTimer(time.Millisecond*time.Duration(3000), now)
	timer.AddTimer(time.Millisecond*time.Duration(3000), now)
	// for i = 0; i < N; i++ {
	// 	timer.AddTimer(time.Millisecond*time.Duration(10*i), now)
	// 	timer.AddTimer(time.Millisecond*time.Duration(10*i), now)
	// }
	fmt.Println("start timer = ", time.Now().Format("2006-01-02 15:04:05"))
	timer.Start()

	fmt.Println("------------分割线-----------------")
	if sum != 2*N {
		t.Error("failed")
	}
}
