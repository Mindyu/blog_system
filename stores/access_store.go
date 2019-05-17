package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetAccessCount(c *gin.Context) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	if err := DB.Debug().Model(&models.Access{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}


func SaveAccess(c *gin.Context, access *models.Access) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(access).Error; err != nil {
		return err
	}
	return nil
}