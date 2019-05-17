package stores

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetNotReadMsgByUserName(c *gin.Context, name string) ([]*models.PrivateMsg, error) {
	msgs := []*models.PrivateMsg{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ? and is_read = ? and receiver = ?", 0, 0, name).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func GetReadMsgByUserName(c *gin.Context, name string) ([]*models.PrivateMsg, error) {
	msgs := []*models.PrivateMsg{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ? and is_read = ? and receiver = ?", 0, 1, name).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func SavePrivateMsg(c *gin.Context, msg *models.PrivateMsg) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(msg).Error; err != nil {
		return err
	}
	return nil
}

func GetPrivateMsgById(c *gin.Context, id int) (*models.PrivateMsg, error) {
	msg := &models.PrivateMsg{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("id = ? and status = ?", id, 0).First(msg).Error; err != nil {
		return nil, err
	}
	return msg, nil
}