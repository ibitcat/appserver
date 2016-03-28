package redis

import (
	"testing"
)

func Test_Connect(t *testing.T) {
	var err error
	rc := Connect()
	rc.Do("SETEX", "test001", 1, 100)

	t.Log("test log : ", err)
}
