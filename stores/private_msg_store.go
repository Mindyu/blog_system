package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetNotReadMsgByUserName(c *gin.Context, name string) ([]*models.PrivateMsg, error) {
	msgs := []*models.PrivateMsg{}

	if err := persistence.GetOrm().Debug().Where("status = ? and is_read = ? and receiver = ?", 0, 0, name).Order("created_at DESC").Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func GetReadMsgByUserName(c *gin.Context, name string) ([]*models.PrivateMsg, error) {
	msgs := []*models.PrivateMsg{}

	if err := persistence.GetOrm().Debug().Where("status = ? and is_read = ? and receiver = ?", 0, 1, name).Order("created_at DESC").Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func SavePrivateMsg(c *gin.Context, msg *models.PrivateMsg) error {

	if err := persistence.GetOrm().Save(msg).Error; err != nil {
		return err
	}
	return nil
}

func GetPrivateMsgById(c *gin.Context, id int) (*models.PrivateMsg, error) {
	msg := &models.PrivateMsg{}

	if err := persistence.GetOrm().Debug().Where("id = ? and status = ?", id, 0).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}