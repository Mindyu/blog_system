package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func EncryptPasswordWithSalt(password, salt string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	md5Ctx.Write([]byte(salt))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
