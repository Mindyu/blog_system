package handlers

import (
	"errors"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func TokenHelper(c *gin.Context, user *models.User) (string, error) {
	// 根据user的权限和角色ID查询
	roleId := user.RoleID

	role, err := stores.GetRoleByIds(c, roleId)
	if err!=nil {
		return "", errors.New("获取角色")
	}

	roleAuthList, err := stores.GetRoleAuthByRoleID(c, roleId)
	if err!=nil {
	 	return "", errors.New("获取角色权限失败")
	}
	authIds := []int{}
	for _, roleAuth := range roleAuthList{
		authIds = append(authIds, roleAuth.AuthID)
	}
	authList, err := stores.GetAuthByIds(c, authIds)
	auths := []string{}
	for _, auth := range authList {
		auths = append(auths, auth.AuthName)
	}

	return utils.GenerateToken(user.Username, role.RoleName, strings.Join(auths, "|"))
}
