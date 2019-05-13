package fileutil

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Mindyu/blog_system/utils/md5"
	"io"
)

// ExtensionName 截取字符串 start 起点下标
func ExtensionName(str string) string {
	rs := []rune(str)
	var start int
	for i := range rs {
		if string(rs[i]) == "." {
			start = i
		}
	}
	return string(rs[start:])
}

// UniqueID 生成Guid字串
func UniqueID() string {
	b := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return md5.EncryptPassword(base64.URLEncoding.EncodeToString(b))
}
