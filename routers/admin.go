package routers

import (
	"github.com/Mindyu/blog_system/utils"
	"github.com/Mindyu/blog_system/utils/systemlog"
	"github.com/Mindyu/blog_system/views"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(router *gin.Engine){
	router.GET("/ws", views.WebSocket)

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", views.Login)
		userRouter.POST("/add", systemlog.OperationLog(views.AddUser, "新增用户"))
		userRouter.PUT("/edit", systemlog.OperationLog(views.UpdateUser, "修改用户信息"))
		userRouter.GET("/query", views.QueryUserByName)
		userRouter.DELETE("/delete", utils.BasicAuth(systemlog.OperationLog(views.DeleteUserById, "删除用户")))
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
		blogRouter.POST("/type/add", systemlog.OperationLog(views.AddBlogType,"新增博客类别"))
		blogRouter.GET("/typecount", views.QueryBlogTypeStats)
		blogRouter.GET("/monthcount", views.QueryBlogByMonth)
		blogRouter.GET("/tags", views.QueryBlogTags)
		blogRouter.POST("/add", systemlog.OperationLog(views.AddBlog, "新增博客"))
		blogRouter.PUT("/update", systemlog.OperationLog(views.UpdateBlog, "修改博客"))
		blogRouter.DELETE("/delete", utils.BasicAuth(systemlog.OperationLog(views.DeleteBlogById, "删除博客")))
	}
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/list", views.GetCommentList)                                               // 获取搜索满足条件的评论
		commentRouter.POST("/blogId", views.GetCommentListByBolgId)                                     // 根据博客ID查询所有满足条件的评论
		commentRouter.DELETE("/delete", systemlog.OperationLog(views.DeleteCommentById, "删除评论"))        // 根据评论ID删除评论
		commentRouter.DELETE("/batchDelete", systemlog.OperationLog(views.BatchDeleteComment, "批量删评论")) // 批量删除评论
		commentRouter.POST("/add", views.InsertComment)                                                 // 新建评论
		//commentRouter.POST("/add", systemlog.OperationLog(views.InsertComment, "新增评论"))                 // 新建评论
	}
	replyRouter := router.Group("/reply")
	{
		replyRouter.POST("/add", systemlog.OperationLog(views.ReplyComment, "新增回复"))                        // 回复
		replyRouter.POST("/list", views.GetReplyList)                                                       // 获取搜索满足条件的回复
		replyRouter.DELETE("/delete", systemlog.OperationLog(views.DeleteCommentReplyById, "删除回复"))         // 根据评论ID删除评论
		replyRouter.DELETE("/batchDelete", systemlog.OperationLog(views.BatchDeleteCommentReply, "批量删除回复")) // 批量删除评论
	}
	friendRouter := router.Group("/friend")
	{
		friendRouter.POST("/add", systemlog.OperationLog(views.AddFriend, "新增好友"))
		friendRouter.POST("/list", views.GetFriendList)
		friendRouter.DELETE("/delete", systemlog.OperationLog(views.DeleteFriendByName, "删除好友关系"))
	}
	attentionRouter := router.Group("/attention")
	{
		attentionRouter.POST("/add", systemlog.OperationLog(views.AddAttention, "新增关注关系"))
		attentionRouter.POST("/list", views.GetAttentionList)
		attentionRouter.DELETE("/delete", systemlog.OperationLog(views.DeleteAttentionByName, "删除关注关系"))
	}
	systemLogRouter := router.Group("/system")
	{
		systemLogRouter.POST("/list", views.GetSystemLogList)
		systemLogRouter.GET("/count", views.GetSystemLogCount)
		systemLogRouter.GET("/access", views.QuerySystemAccessCount)
		systemLogRouter.GET("/access/week", views.GetSystemAccessWeek)
		systemLogRouter.DELETE("/delete", systemlog.OperationLog(views.DeleteSystemLogById, "删除日志记录"))
		systemLogRouter.GET("/export", systemlog.OperationLog(views.ExportOperationLog, "导出操作日志"))
	}
	privateMsgRouter := router.Group("/msg")
	{
		privateMsgRouter.GET("/unread", views.QueryNotReadMsgByName)
		privateMsgRouter.GET("/read", views.QueryReadMsgByName)
		privateMsgRouter.PUT("/read", systemlog.OperationLog(views.ReadPrivateMsg, "已阅私信"))
		privateMsgRouter.POST("/add", systemlog.OperationLog(views.AddPrivateMsg, "发送私信"))
	}

	router.POST("/file/upload", systemlog.OperationLog(views.Upload, "上传文件"))
}
