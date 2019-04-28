package utils

import (
	"github.com/Mindyu/blog_system/middleware/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

const base64Secret = "MDk8ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzmjYyN2I0ZjY="

func GenerateToken(username string, roles string, auths string) (string, error) {
	j := &jwt.JWT{
		SigningKey: []byte(base64Secret),
	}
	claims := jwt.CustomClaims{
		UserName:username,
		UserRole:roles,
		UserAuth: auths,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "mindyu",                   		//签名的发行者
		},
	}

	return j.CreateToken(claims)
}

