package stores

import (
	"fmt"
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
	if err := DB.Debug().Where("comment_id = ? and status = ?", commentId, 0).Order("updated_at DESC").
		Find(&commentReply).Error; err != nil {
		return nil, err
	}
	return commentReply, nil
}

func GetCommentReplyList(c *gin.Context, page, pageSize, commentId int, searchKey string) ([]*models.CommentReply, error) {
	replys := []*models.CommentReply{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if commentId != 0 {
		sql = fmt.Sprintf("%s and comment_id = %d", sql, commentId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (from_username LIKE '%%%s%%') or (to_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order("created_at DESC").Find(&replys).Error; err != nil {
		return nil, err
	}
	return replys, nil
}

func GetCommentReplyListCount(c *gin.Context, commentId int, searchKey string) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if commentId != 0 {
		sql = fmt.Sprintf("%s and comment_id = %d", sql, commentId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (from_username LIKE '%%%s%%') or (to_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Model(&models.CommentReply{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetCommentReplyByID(c *gin.Context, commentId int) (*models.CommentReply, error) {
	reply := &models.CommentReply{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ?", 0).First(reply, commentId).Error; err != nil {
		return nil, err
	}
	return reply, nil
}

func GetCommentReplyByIDs(c *gin.Context, replyIds []int) ([]*models.CommentReply, error) {
	replys := []*models.CommentReply{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ? and id in (?)", 0, replyIds).Find(&replys).Error; err != nil {
		return nil, err
	}
	return replys, nil
}

func GetCommentReplyListWithAuthor(c *gin.Context, page, pageSize, commentId int, name, searchKey string) ([]*models.CommentReply, error) {
	replys := []*models.CommentReply{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if commentId != 0 {
		sql = fmt.Sprintf("%s and comment_id = %d", sql, commentId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (from_username LIKE '%%%s%%') or (to_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}

	originSql := fmt.Sprintf(`SELECT 
	reply.id,comment_id,reply_id,reply_type,from_username,to_username,status,created_at,updated_at
FROM comment_reply reply
LEFT JOIN 
(
	SELECT id
	from comment
	WHERE status = 0 and blog_author = '%s'
)com
on 
reply.comment_id = com.id
WHERE 1=1 and com.id is not null and %s
ORDER BY
updated_at DESC
LIMIT %d OFFSET %d`, name, sql, pageSize, (page-1)*pageSize)
	if err := DB.Debug().Raw(originSql).Scan(&replys).Error; err != nil {
		return nil, err
	}
	return replys, nil
}

func GetCommentReplyListCountWithAuthor(c *gin.Context, commentId int, name, searchKey string) (int, error) {
	count := struct {
		Count int
	}{Count: 0}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if commentId != 0 {
		sql = fmt.Sprintf("%s and comment_id = %d", sql, commentId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (from_username LIKE '%%%s%%') or (to_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	originSql := fmt.Sprintf(`SELECT 
	count(*) as count
FROM comment_reply reply
LEFT JOIN 
(
	SELECT id, blog_author
	from comment
	WHERE status = 0 and blog_author = '%s'
)com
on 
reply.comment_id = com.id
WHERE 1=1 and com.id is not null and %s`, name, sql)
	if err := DB.Debug().Raw(originSql).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count.Count, nil
}
