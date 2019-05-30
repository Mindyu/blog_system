package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetRoleById(c *gin.Context, roleId int) (*models.Role, error) {
	role := &models.Role{}

	if err := persistence.GetOrm().Debug().Where("id = ?", roleId).Find(role).Error; err!=nil{
		return nil, err
	}
	return role, nil
}


func GetAllRoles(c *gin.Context) ([]*models.Role, error) {
	roles := []*models.Role{}

	if err := persistence.GetOrm().Find(&roles).Error; err!=nil{
		return nil, err
	}
	return roles, nil
}


func SaveRole(c *gin.Context, role *models.Role) error {

	if err := persistence.GetOrm().Save(role).Error; err != nil {
		return err
	}
	return nil
}