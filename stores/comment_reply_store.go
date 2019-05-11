package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func SaveCommentReply(c *gin.Context, comment *models.CommentReply) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(comment).Error; err != nil {
		return err
	}
	return nil
}

func QueryCommentReplyByCommentID(c *gin.Context, commentId int) ([]*models.CommentReply, error) {
	commentReply := []*models.CommentReply{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("comment_id = ? and status = ?", commentId, 0).Find(&commentReply).Error; err != nil {
		return nil, err
	}
	return commentReply, nil
}
