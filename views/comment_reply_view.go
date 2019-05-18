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

func GetReplyList(c *gin.Context) {
	param := &common.ReplyPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)
	name := utils.InjectUserName(c)

	if name == "" {
		comments, err := stores.GetCommentReplyList(c, param.CurrentPage, param.PageSize, param.CommentId, param.SearchWords)
		if err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}
		total, err := stores.GetCommentReplyListCount(c, param.CommentId, param.SearchWords)
		if err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}

		utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: comments})
		return
	}
	comments, err := stores.GetCommentReplyListWithAuthor(c, param.CurrentPage, param.PageSize, param.CommentId, name ,param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetCommentReplyListCountWithAuthor(c, param.CommentId, name, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: comments})
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

func DeleteCommentReplyById(c *gin.Context) {
	id := c.Query("replyId")

	commentId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "回复ID转换失败")
		return
	}

	reply, err := stores.GetCommentReplyByID(c, commentId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	reply.Status = 1
	if err := stores.SaveCommentReply(c, reply); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func BatchDeleteCommentReply(c *gin.Context) {
	idStr := c.Query("replyIds")

	ids := strings.Split(idStr, ",")
	replyIds := []int{}
	for _, id := range ids {
		replyId, err := strconv.Atoi(id)
		if err != nil {
			utils.MakeErrResponse(c, "用户ID转换失败")
			return
		}
		replyIds = append(replyIds, replyId)
	}

	replys, err := stores.GetCommentReplyByIDs(c, replyIds)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	for _, replys := range replys {
		replys.Status = 1
		if err := stores.SaveCommentReply(c, replys); err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}
	}
	utils.MakeOkResponse(c, "删除成功")
}
