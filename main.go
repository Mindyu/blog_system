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
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/list", views.GetCommentList)              // 获取搜索满足条件的评论
		commentRouter.POST("/blogId", views.GetCommentListByBolgId)    // 根据博客ID查询所有满足条件的评论
		commentRouter.DELETE("/delete", views.DeleteCommentById)       // 根据评论ID删除评论
		commentRouter.DELETE("/batchDelete", views.BatchDeleteComment) // 批量删除评论
		commentRouter.POST("/add", views.InsertComment)                // 批量删除评论
		commentRouter.POST("/reply", views.ReplyComment)               // 批量删除评论
	}
	friendRouter := router.Group("/friend")
	{
		friendRouter.POST("/list", views.GetFriendList)
		friendRouter.DELETE("/delete", views.DeleteFriendByName)
	}

	router.Run(":8081")
}
