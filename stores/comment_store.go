package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetCommentList(c *gin.Context, page, pageSize, blogId int, searchKey string) ([]*models.Comment, error) {
	comments := []*models.Comment{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if blogId != 0 {
		sql = fmt.Sprintf("%s and blog_id = %d", sql, blogId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (blog_title LIKE '%%%s%%') or (comment_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Where(sql).Offset((page-1)*pageSize).Limit(pageSize).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}


func GetCommentListCount(c *gin.Context, blogId int, searchKey string) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if blogId != 0 {
		sql = fmt.Sprintf("%s and blog_id = %d", sql, blogId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (blog_title LIKE '%%%s%%') or (comment_username LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Model(&models.Comment{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetCommentByBlogID(c *gin.Context, blogId int) ([]*models.Comment, error) {
	comments := []*models.Comment{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("blog_id = ? and status = ?", blogId, 0).Find(comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func GetCommentByID(c *gin.Context, commentId int) (*models.Comment, error) {
	comment := &models.Comment{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ?", 0).First(comment, commentId).Error; err != nil {
		return nil, err
	}
	return comment, nil
}


func GetCommentByIDs(c *gin.Context, commentIds []int) ([]*models.Comment, error) {
	comment := []*models.Comment{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ? and id in (?)", 0, commentIds).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func SaveComment(c *gin.Context, comment *models.Comment) error {
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
