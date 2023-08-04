package utils

import (
	"math/rand"
	"time"
)

// RandString 生成随机字符串作为锁唯一标识
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
