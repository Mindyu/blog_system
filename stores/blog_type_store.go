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

func SaveBlogType(c *gin.Context, blogType *models.BlogType) error {

	if err := persistence.GetOrm().Save(blogType).Error; err != nil {
		return err
	}
	return nil
}