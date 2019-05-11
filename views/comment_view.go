package views

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
	"time"
)

func GetCommentList(c *gin.Context) {
	param := &common.CommentPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	comments, err := stores.GetCommentList(c, param.CurrentPage, param.PageSize, param.BlogId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetCommentListCount(c, param.BlogId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: comments})
}

func GetCommentListByBolgId(c *gin.Context) {
	param := &common.CommentPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		log.Debug(err.Error())
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	comments, err := stores.GetCommentList(c, param.CurrentPage, param.PageSize, param.BlogId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetCommentListCount(c, param.BlogId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	// 评论回复
	for _, comment := range comments {
		replys, _ := stores.QueryCommentReplyByCommentID(c, comment.ID)
		comment.CommentReply = replys
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: comments})

}

func DeleteCommentByBlogId(c *gin.Context) {
	id := c.Query("blogId")

	blogId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "博客ID转换失败")
		return
	}

	comments, err := stores.GetCommentByBlogID(c, blogId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	for _, comment := range comments {
		comment.Status = 1
		if err := stores.SaveComment(c, comment); err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}

	}

	utils.MakeOkResponse(c, "删除成功")
}

func DeleteCommentById(c *gin.Context) {
	id := c.Query("commentId")

	commentId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "评论ID转换失败")
		return
	}

	comment, err := stores.GetCommentByID(c, commentId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	comment.Status = 1
	if err := stores.SaveComment(c, comment); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func BatchDeleteComment(c *gin.Context) {
	idStr := c.Query("commentIds")

	ids := strings.Split(idStr, ",")
	commentIds := []int{}
	for _, id := range ids {
		commentId, err := strconv.Atoi(id)
		if err != nil {
			utils.MakeErrResponse(c, "用户ID转换失败")
			return
		}
		commentIds = append(commentIds, commentId)
	}

	comments, err := stores.GetCommentByIDs(c, commentIds)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	for _, comment := range comments {
		comment.Status = 1
		if err := stores.SaveComment(c, comment); err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}
	}
	utils.MakeOkResponse(c, "删除成功")
}

func InsertComment(c *gin.Context) {
	comment := &models.Comment{}

	err := c.ShouldBindJSON(comment)
	if err != nil {
		utils.MakeErrResponse(c, "评论信息转换失败")
		return
	}

	comment.Status = 0
	comment.CreatedAt = time.Now()

	if err := stores.SaveComment(c, comment); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "评论成功")
}

func ReplyComment(c *gin.Context) {
	commentReply := &models.CommentReply{}

	err := c.ShouldBindJSON(commentReply)
	if err != nil {
		utils.MakeErrResponse(c, "回复评论信息转换失败")
		return
	}

	commentReply.Status = 0
	commentReply.CreatedAt = time.Now()

	if err := stores.SaveCommentReply(c, commentReply); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "回复成功")
}
