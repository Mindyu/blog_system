package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetAuthByIds(c *gin.Context, authIds []int) ([]*models.Auth, error) {
	auths := []*models.Auth{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("id in (?)", authIds).Find(&auths).Error; err!=nil{
		return nil, err
	}
	return auths, nil
}

