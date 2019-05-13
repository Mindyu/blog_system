package utils

import (
	"math/rand"
	"time"
)

// return len=32  salt
func GetRandomSalt() string {
	return GetRandomString(32)
}

//生成随机字符串
func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
