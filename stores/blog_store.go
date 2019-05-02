package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetBlogList(c *gin.Context, page, pageSize, blogTypeId int, searchKey string) ([]*models.Blog, error) {
	blogs := []*models.Blog{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if blogTypeId != 0 {
		sql = fmt.Sprintf("%s and type_id = %d", sql, blogTypeId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (blog_title LIKE '%%%s%%') or (blog_content LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Where(sql).Offset((page-1)*pageSize).Limit(pageSize).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}


func GetBlogListCount(c *gin.Context, blogTypeId int, searchKey string) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if blogTypeId != 0 {
		sql = fmt.Sprintf("%s and type_id = %d", sql, blogTypeId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (blog_title LIKE '%%%s%%') or (blog_content LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Model(&models.Blog{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetBlogById(c *gin.Context, id int) (*models.Blog, error) {
	blog := &models.Blog{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("id = ? and status = ?", id, 0).First(blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}


func SaveBlog(c *gin.Context, blog *models.Blog) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(blog).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBlogById(c *gin.Context, user *models.User) error{
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}