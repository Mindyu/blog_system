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
		userRouter.GET("/query", views.QueryUserById)
		userRouter.DELETE("/delete", views.DeleteUserById)
		userRouter.GET("/valid/:name", views.ValidUserName)
		userRouter.GET("/auth", views.QueryUserAuth)
		userRouter.GET("/type", views.QueryUserType)
		userRouter.POST("/all", views.QueryAllUser)
	}

	blogRouter := router.Group("/blog")
	{
		blogRouter.POST("/list", views.GetBlogList)
		blogRouter.GET("/query", views.QueryBlogById)
		blogRouter.GET("/type", views.QueryAllBlogType)
		blogRouter.POST("/add", views.AddBlog)
		blogRouter.PUT("/update", views.UpdateBlog)
		blogRouter.DELETE("/delete", views.DeleteBlogById)
	}

	router.Run(":8081")
}
