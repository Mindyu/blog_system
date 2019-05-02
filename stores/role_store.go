package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetRoleByIds(c *gin.Context, roleId int) (*models.Role, error) {
	role := &models.Role{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("id = ?", roleId).Find(role).Error; err!=nil{
		return nil, err
	}
	return role, nil
}


func GetAllRoles(c *gin.Context) ([]*models.Role, error) {
	roles := []*models.Role{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Find(&roles).Error; err!=nil{
		return nil, err
	}
	return roles, nil
}