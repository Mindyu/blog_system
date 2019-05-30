package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
)

func GetAttentionList(c *gin.Context, page, pageSize int, username string) ([]*models.Attention, error) {
	attentions := []*models.Attention{}

	sql := fmt.Sprintf("status = %d", 0)
	if username != "" {
		sql = fmt.Sprintf("%s and focus_user = '%s'", sql, username)
	}
	if err := persistence.GetOrm().Debug().Where(sql).Offset((page - 1) * pageSize).Limit(pageSize).Order("updated_at DESC").Find(&attentions).Error; err != nil {
		return nil, err
	}
	return attentions, nil
}

func GetAttentionListCount(c *gin.Context, username string) (int, error) {
	count := 0

	sql := fmt.Sprintf("status = %d", 0)
	if username != "" {
		sql = fmt.Sprintf("%s and focus_user = '%s'", sql, username)
	}
	if err := persistence.GetOrm().Debug().Model(&models.Attention{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetAttentionByID(c *gin.Context, id int) (*models.Attention, error) {
	attention := &models.Attention{}

	if err := persistence.GetOrm().Debug().Where("status = ?", 0).First(attention, id).Error; err != nil {
		return nil, err
	}
	return attention, nil
}

func GetAttentionByName(c *gin.Context, userName, attentionName string) (*models.Attention, error) {
	attention := &models.Attention{}

	if err := persistence.GetOrm().Debug().Where("status = ? and (focus_user = ? and focused_user = ?)",
		0, userName, attentionName).First(attention).Error; err != nil {
		return nil, err
	}
	return attention, nil
}

func SaveAttention(c *gin.Context, attention *models.Attention) error {

	if err := persistence.GetOrm().Save(attention).Error; err != nil {
		return err
	}
	return nil
}
