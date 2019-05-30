package systemlog

import (
	"bytes"
	"github.com/Mindyu/blog_system/middleware/jwt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"time"
)

func OperationLog(h gin.HandlerFunc, operation string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ""
		if ctx.Request.URL.Path != "/user/add" {
			claims := ctx.MustGet("claims")
			customclaims := claims.(*jwt.CustomClaims)
			userName = customclaims.UserName
		}
		data, _ := ctx.GetRawData()

		systemLog := &models.Log{}
		if ctx.Request.URL.Path != "/file/upload" {
			systemLog.Params = string(data)
			if systemLog.Params == "" {
				systemLog.Params = ctx.Request.URL.RawQuery
			}
		}
		systemLog.Username = userName
		systemLog.CreatedAt = time.Now()
		systemLog.CallAPI = ctx.Request.URL.Path
		systemLog.Operation = operation
		systemLog.Status = 0
		log.Info(systemLog)
		_ = stores.SaveSystemLog(ctx, systemLog)

		// 将body值写回
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		h(ctx)
	}
}
