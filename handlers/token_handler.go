package handlers

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func TokenHelper(c *gin.Context, user *models.User) (string, error) {
	// 根据user的权限和角色ID查询

	return utils.GenerateToken(user.Username, "admin", "edit|add")
}
