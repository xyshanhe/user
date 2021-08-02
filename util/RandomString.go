package util

import (
	"math/rand"
	"time"
)

// RandomString 用户随机名称
func RandomString(n int) string {
	//名称没有填写的时候，系统随机生成一个名称
	var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
