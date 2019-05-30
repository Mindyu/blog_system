package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetAllBlogType(c *gin.Context) ([]*models.BlogType, error) {
	types := []*models.BlogType{}

	if err := persistence.GetOrm().Find(&types).Error; err!=nil{
		return nil, err
	}
	return types, nil
}