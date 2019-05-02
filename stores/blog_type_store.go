package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetAllBlogType(c *gin.Context) ([]*models.BlogType, error) {
	types := []*models.BlogType{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Find(&types).Error; err!=nil{
		return nil, err
	}
	return types, nil
}