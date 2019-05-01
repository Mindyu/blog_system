package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context, userId int) (*models.User, error) {
	user := &models.User{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Where("status = ?", 0).First(user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func GetUserByName(c *gin.Context, userName string) (*models.User, error) {
	user := &models.User{}
	DB, err := utils.InitDB()
	//defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("username = ? and status = ?", userName, 0).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUser(c *gin.Context, user *models.User) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func Delete(c *gin.Context, user *models.User) error{
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}