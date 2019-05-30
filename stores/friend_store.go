package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetFriendList(c *gin.Context, page, pageSize int, username string) ([]*models.Friend, error) {
	friends := []*models.Friend{}

	sql := fmt.Sprintf("status = %d", 0)
	if username != "" {
		sql = fmt.Sprintf("%s and ((username_1 = '%s') or (username_2 = '%s'))", sql, username, username)
	}
	if err := persistence.GetOrm().Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order("updated_at DESC").Find(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func GetFriendListCount(c *gin.Context, username string) (int, error) {
	count := 0

	sql := fmt.Sprintf("status = %d", 0)
	if username != "" {
		sql = fmt.Sprintf("%s and ((username_1 = '%s') or (username_2 = '%s'))", sql, username, username)
	}
	if err := persistence.GetOrm().Debug().Model(&models.Friend{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetFriendByID(c *gin.Context, id int) (*models.Friend, error) {
	friend := &models.Friend{}

	if err := persistence.GetOrm().Debug().Where("status = ?", 0).First(friend, id).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func GetFriendByName(c *gin.Context, userName, friendName string) (*models.Friend, error) {
	friend := &models.Friend{}

	if err := persistence.GetOrm().Debug().Where("status = ? and ((username_1 = ? and username_2 = ?) or (username_1 = ? and username_2 = ?))",
		0, userName, friendName, friendName, userName).First(friend).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func SaveFriend(c *gin.Context, friend *models.Friend) error {

	if err := persistence.GetOrm().Save(friend).Error; err != nil {
		return err
	}
	return nil
}
