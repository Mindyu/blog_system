package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetRoleAuthByRoleID(c *gin.Context, roleId int) ([]*models.RoleAuth, error) {
	roleAuth := []*models.RoleAuth{}

	if err := persistence.GetOrm().Debug().Where("role_id = ?", roleId).Find(&roleAuth).Error; err != nil {
		return nil, err
	}
	return roleAuth, nil
}


