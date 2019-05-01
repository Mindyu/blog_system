package handlers

import (
	"errors"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils/md5"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, userName string, password string) (*models.User, error) {

	user, err := stores.GetUserByName(c, userName)
	if err!=nil {
		return nil, errors.New("用户名不存在")
	}
	if md5.EncryptPassword(password) == user.Password {
		return user, nil
	}
	return nil, errors.New("密码错误")
}
