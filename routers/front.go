package routers

import (
	"github.com/Mindyu/blog_system/views"
	"github.com/gin-gonic/gin"
)

func NewFrontRouter(engine *gin.Engine){

	frontRouter := engine.Group("/front")
	{
		frontRouter.POST("/blog/list", views.GetBlogList)
		frontRouter.GET("/blog/query", views.QueryBlogById)
		frontRouter.GET("/blog/typecount", views.QueryBlogTypeStats)
		frontRouter.GET("/blog/monthcount", views.QueryBlogByMonth)
		frontRouter.GET("/blog/tags", views.QueryBlogTags)
		frontRouter.GET("/blog/sug", views.GetBlogSearchKeySug)

		frontRouter.POST("/comment/blogId", views.GetCommentListByBolgId)   // 根据博客ID查询所有满足条件的评论
		frontRouter.POST("/comment/add", views.InsertComment)

		frontRouter.POST("/reply/add", views.ReplyComment)   // 回复
	}

}