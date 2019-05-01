package main

import (
	"github.com/Mindyu/blog_system/middleware"
	"github.com/Mindyu/blog_system/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//使用中间件
	router.Use(middleware.Cors())

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", views.Login)
		userRouter.POST("/add", views.AddUser)
		userRouter.PUT("/edit", views.UpdateUser)
		userRouter.GET("/query", views.GetUser)
		userRouter.DELETE("/delete", views.DeleteUser)
	}

	router.Run(":8081")
}
