package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetRoleAuthByRoleID(c *gin.Context, roleId int) ([]*models.RoleAuth, error) {
	roleAuth := []*models.RoleAuth{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err = DB.Debug().Where("role_id = ?", roleId).Find(&roleAuth).Error; err != nil {
		return nil, err
	}
	return roleAuth, nil
}


