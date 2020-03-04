package util

import "math/rand"

// 返回随机的字符串
func GenRandomStr(n int) string {
	letterByte := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letterByte[rand.Intn(len(letterByte))]
	}

	return string(b)
}
