package stores

import (
	"fmt"
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
	defer DB.Close()
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

func DeleteUser(c *gin.Context, user *models.User) error{
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

func GetUserList(c *gin.Context, page, pageSize, roleId int, searchKey string) ([]*models.User, error) {
	user := []*models.User{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if roleId != 0 {
		sql = fmt.Sprintf("%s and role_id = %d", sql, roleId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username LIKE '%%%s%%') or (nickname LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Where(sql).Offset((page-1)*pageSize).Limit(pageSize).Order("created_at DESC").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}


func GetUserListCount(c *gin.Context, roleId int, searchKey string) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if roleId != 0 {
		sql = fmt.Sprintf("%s and role_id = %d", sql, roleId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username LIKE '%%%s%%') or (nickname LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Model(&models.User{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}