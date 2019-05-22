package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetBlogList(c *gin.Context, page, pageSize, blogTypeId int, searchKey, author string, sortType int) ([]*models.Blog, error) {
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
		if len(searchKey) == 7 && strings.Index(searchKey, "-") == 4 { // 按日期搜索
			sql = fmt.Sprintf("%s and created_at like '%s%%'", sql, searchKey)
		} else {
			sql = fmt.Sprintf("%s and ((blog_title LIKE '%%%s%%') or (keywords LIKE '%%%s%%') or "+
				"(author LIKE '%%%s%%'))", sql, searchKey, searchKey, searchKey)
		}
	}
	if author!="" {
		sql = fmt.Sprintf("%s and author = '%s'", sql, author)
	}
	sortList := []string{"updated_at", "read_count", "reply_count"}
	if err := DB.Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order(sortList[sortType] + " DESC").
		Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlogListCount(c *gin.Context, blogTypeId int, searchKey, author string) (int, error) {
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
	if author!="" {
		sql = fmt.Sprintf("%s and author = '%s'", sql, author)
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

func DeleteBlogById(c *gin.Context, user *models.User) error {
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

func GetBlogTypeStats(c *gin.Context) ([]*common.BlogTypeResp, error) {
	types := []*common.BlogTypeResp{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := `SELECT 
	type_id, type_name, count(*) as count
from 
	blog 
LEFT JOIN
	blog_type type on type_id = type.id
WHERE
	blog.status = 0
GROUP BY
	type_id
ORDER BY
	count DESC`
	sql = strings.Replace(sql, "\r\n", "\n", -1)
	if err := DB.Debug().Raw(sql).Scan(&types).Error; err != nil {
		return nil, err
	}
	return types, nil
}

type Stats struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func GetBlogStatsByMonth(c *gin.Context) ([]*Stats, error) {
	stats := []*Stats{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := `SELECT 
	DATE_FORMAT(created_at, '%Y-%m') as month, count(*) as count
from 
	blog
GROUP BY
	month
ORDER BY
	month DESC`
	sql = strings.Replace(sql, "\r\n", "\n", -1)
	if err := DB.Debug().Raw(sql).Scan(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

func GetBlogTags(c *gin.Context) ([]*common.Key, error) {
	tags := []*common.Key{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := `SELECT 
	keywords
from 
	blog
WHERE
	status = 0`
	sql = strings.Replace(sql, "\r\n", "\n", -1)
	if err := DB.Debug().Raw(sql).Scan(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func SaveBlogType(c *gin.Context, blogType *models.BlogType) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(blogType).Error; err != nil {
		return err
	}
	return nil
}