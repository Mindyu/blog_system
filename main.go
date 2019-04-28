package main

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//使用中间件
	router.Use(Cors())

	router.POST("/jwt/token", token)

	router.Run(":8081")
}

func token(c *gin.Context){
	username := c.Request.PostFormValue("username")
	password := c.Request.PostFormValue("password")

	if username != "zhangchunhui" {
		utils.MakeOkResponse(c, "用户名不存在")
	}
	if password != "123456" {
		utils.MakeOkResponse(c, "密码错误")
	}

	accessToken, err := utils.GenerateToken(username, "admin", "edit|add")
	if err!=nil {
		utils.MakeErrResponse(c, "", "获取token失败")
	}

	accessTokenObj := models.AccessToken{
		Access_token:accessToken,
		Token_type:"bearer",
	}
	utils.MakeOkResponse(c, accessTokenObj)
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, x-requested-with, Authorization,userID,Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}