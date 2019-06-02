package jwt

import (
	"errors"
	"github.com/Mindyu/blog_system/config"
	"github.com/Mindyu/blog_system/utils/set"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var permitSet *set.Set
var SignKey string

func init() {
	permitSet = set.New()
	permitSet.Add("/user/login")
	permitSet.Add("/user/add")
	// permitSet.Add("/file/upload")    // 对文件上传操作应该加日志
	SignKey = config.Config().SecretKey
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		permitSet.Has(c.Request.URL.Path)
		if permitSet.Has(c.Request.URL.Path) || strings.Contains(c.Request.URL.Path, "/user/valid") {
			c.Next()
			return
		}
		token := c.Request.Header.Get("Authorization")
		if c.Request.URL.Path == "/file/upload" || c.Request.URL.Path == "/ws" { // 对于文件上传时，则通过拼接token的方式
			token = c.Query("token")
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, "请求未携带token，无权限访问")
			c.Abort()
			return
		}

		log.Print("get token: ", token)

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, "授权已过期")
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		log.Println(claims)
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		// before request
		c.Next()
		// after request
		return
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserName string `json:"user_name"`
	UserRole string `json:"user_role"`
	UserAuth string `json:"user_auth"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
