package main

import (
	"github.com/Mindyu/blog_system/middleware"
	"github.com/Mindyu/blog_system/middleware/jwt"
	"github.com/Mindyu/blog_system/utils"
	"github.com/Mindyu/blog_system/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.StaticFS("/file", http.Dir("public/upload"))
	//使用中间件
	router.Use(middleware.Cors())   // 跨域请求解决
	router.Use(jwt.JWTAuth())       // Jwt认证，除登陆外所有请求都需要携带tokenren认证

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", views.Login)
		userRouter.POST("/add", views.AddUser)
		userRouter.PUT("/edit", views.UpdateUser)
		userRouter.GET("/query", views.QueryUserById)
		userRouter.DELETE("/delete", utils.BasicAuth(views.DeleteUserById))
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
		blogRouter.DELETE("/delete", utils.BasicAuth(views.DeleteBlogById))
	}
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/list", views.GetCommentList)              // 获取搜索满足条件的评论
		commentRouter.POST("/blogId", views.GetCommentListByBolgId)    // 根据博客ID查询所有满足条件的评论
		commentRouter.DELETE("/delete", views.DeleteCommentById)       // 根据评论ID删除评论
		commentRouter.DELETE("/batchDelete", views.BatchDeleteComment) // 批量删除评论
		commentRouter.POST("/add", views.InsertComment)                // 新建评论
	}
	replyRouter := router.Group("/reply")
	{
		replyRouter.POST("/add", views.ReplyComment)                    // 回复
		replyRouter.POST("/list", views.GetReplyList)                 // 获取搜索满足条件的回复
		replyRouter.DELETE("/delete", views.DeleteCommentReplyById)       // 根据评论ID删除评论
		replyRouter.DELETE("/batchDelete", views.BatchDeleteCommentReply) // 批量删除评论
	}
	friendRouter := router.Group("/friend")
	{
		friendRouter.POST("/list", views.GetFriendList)
		friendRouter.DELETE("/delete", views.DeleteFriendByName)
	}

	router.POST("/file/upload", views.Upload)

	router.Run(":8081")
}
