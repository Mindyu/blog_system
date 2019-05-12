package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetFriendList(c *gin.Context, page, pageSize int, searchKey string) ([]*models.Friend, error) {
	friends := []*models.Friend{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username_1 = '%s') or (username_2 = '%s')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order("updated_at DESC").Find(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func GetFriendListCount(c *gin.Context, searchKey string) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username_1 = '%s') or (username_2 = '%s')", sql, searchKey, searchKey)
	}
	if err := DB.Debug().Model(&models.Friend{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetFriendByID(c *gin.Context, id int) (*models.Friend, error) {
	friend := &models.Friend{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ?", 0).First(friend, id).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func GetFriendByName(c *gin.Context, userName, friendName string) (*models.Friend, error) {
	friend := &models.Friend{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ? and ((username_1 = ? and username_2 = ?) or (username_1 = ? and username_2 = ?))",
		0, userName, friendName, friendName, userName).First(friend).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func SaveFriend(c *gin.Context, friend *models.Friend) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(friend).Error; err != nil {
		return err
	}
	return nil
}