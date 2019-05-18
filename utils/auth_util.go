package utils

import (
	"github.com/Mindyu/blog_system/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

/*var memo *cache.Memo

func init() {
	memo = cache.New(roleGetFunction)
	return
}

// 根据角色获取所有权限
func roleGetFunction(roleid int) ([]string, error) {
	funclist, err := access.GetFunctionNames(roleid)
	return funclist, err
}*/

// BasicAuth 是登录认证，用户分权限管理
func BasicAuth(h gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, exist := ctx.Get("claims")
		if !exist {
			MakeErrResponse(ctx, "获取认证失败")
			return
		}
		customclaims := claims.(*jwt.CustomClaims)
		if customclaims == nil {
			MakeErrResponse(ctx, "解析认证失败")
			return
		}
		if customclaims.UserRole == "admin" || customclaims.UserRole == "superadmin" {
			h(ctx)
		} else {
			MakeErrResponse(ctx, "对不起，您还没有权限")
		}
	}
}

func GetUserNameAndRole(ctx *gin.Context) (name, role string) {
	claims, exist := ctx.Get("claims")
	if !exist {
		log.Info("获取用户名和角色失败")
		return
	}
	customclaims := claims.(*jwt.CustomClaims)
	name = customclaims.UserName
	role = customclaims.UserRole
	return
}

func InjectUserName(ctx *gin.Context) string {
	name, role := GetUserNameAndRole(ctx)
	if role != "admin" && role != "superadmin" {
		return name
	}
	return ""
}
