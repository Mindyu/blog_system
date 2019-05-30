package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context, userId int) (*models.User, error) {
	user := &models.User{}

	if err := persistence.GetOrm().Where("status = ?", 0).First(user, userId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(c *gin.Context, userName string) (*models.User, error) {
	user := &models.User{}

	if err := persistence.GetOrm().Debug().Where("username = ? and status = ?", userName, 0).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsersByNames(c *gin.Context, userNames []string) ([]*models.User, error) {
	users := []*models.User{}

	if err := persistence.GetOrm().Debug().Where("username in (?) and status = ?", userNames, 0).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersByNamesAndLikeWord(c *gin.Context, userNames []string, keyword string) ([]*models.User, error) {
	users := []*models.User{}

	var err error
	if keyword != "" {
		likeStr := fmt.Sprintf("%%%s%%", keyword)
		err = persistence.GetOrm().Debug().Where("username in (?) and status = ? and (username like ? or nickname like ?)",
			userNames, 0, likeStr, likeStr).Find(&users).Error
	} else {
		err = persistence.GetOrm().Debug().Where("username in (?) and status = ?",
			userNames, 0).Find(&users).Error
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUser(c *gin.Context, user *models.User) error {

	if err := persistence.GetOrm().Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(c *gin.Context, user *models.User) error {

	if err := persistence.GetOrm().Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserList(c *gin.Context, page, pageSize, roleId int, searchKey string) ([]*models.User, error) {
	user := []*models.User{}

	sql := fmt.Sprintf("status = %d", 0)
	if roleId != 0 {
		sql = fmt.Sprintf("%s and role_id = %d", sql, roleId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username LIKE '%%%s%%') or (nickname LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := persistence.GetOrm().Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order("created_at DESC").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserListCount(c *gin.Context, roleId int, searchKey string) (int, error) {
	count := 0

	sql := fmt.Sprintf("status = %d", 0)
	if roleId != 0 {
		sql = fmt.Sprintf("%s and role_id = %d", sql, roleId)
	}
	if searchKey != "" {
		sql = fmt.Sprintf("%s and (username LIKE '%%%s%%') or (nickname LIKE '%%%s%%')", sql, searchKey, searchKey)
	}
	if err := persistence.GetOrm().Debug().Model(&models.User{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
