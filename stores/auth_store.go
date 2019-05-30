package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetAuthByIds(c *gin.Context, authIds []int) ([]*models.Auth, error) {
	auths := []*models.Auth{}

	if err := persistence.GetOrm().Debug().Where("id in (?)", authIds).Find(&auths).Error; err!=nil{
		return nil, err
	}
	return auths, nil
}

