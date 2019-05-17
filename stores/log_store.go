package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func GetSystemLogList(c *gin.Context, req common.LogPageRequest) ([]*models.Log, error) {
	systemLogs := []*models.Log{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if req.UserName != "" {
		sql = fmt.Sprintf("%s and call_name like '%s'", sql, req.UserName)
	}
	if req.CallApi != "" {
		sql = fmt.Sprintf("%s and call_api like '%s'", sql, req.CallApi)
	}
	if req.StartTime != "" {
		sql = fmt.Sprintf("%s and created_at > '%s'", sql, req.StartTime)
	}
	if req.EndTime != "" {
		sql = fmt.Sprintf("%s and created_at <= '%s'", sql, req.EndTime)
	}
	if err := DB.Debug().Where(sql).Offset((req.CurrentPage - 1) * req.PageSize).Limit(req.PageSize).
		Order("updated_at DESC").Find(&systemLogs).Error; err != nil {
		return nil, err
	}
	return systemLogs, nil
}

func GetSystemLogListCount(c *gin.Context, req common.LogPageRequest) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if req.UserName != "" {
		sql = fmt.Sprintf("%s and call_name like '%s'", sql, req.UserName)
	}
	if req.CallApi != "" {
		sql = fmt.Sprintf("%s and call_api like '%s'", sql, req.CallApi)
	}
	if req.StartTime != "" {
		sql = fmt.Sprintf("%s and created_at > '%s'", sql, req.StartTime)
	}
	if req.EndTime != "" {
		sql = fmt.Sprintf("%s and created_at <= '%s'", sql, req.EndTime)
	}
	if err := DB.Debug().Model(&models.Log{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}


func GetSystemLogByID(c *gin.Context, id int) (*models.Log, error) {
	systemLog := &models.Log{}
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	if err := DB.Debug().Where("status = ?", 0).First(systemLog, id).Error; err != nil {
		return nil, err
	}
	return systemLog, nil
}


func SaveSystemLog(c *gin.Context, systemLog *models.Log) error {
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return err
	}
	if err := DB.Save(systemLog).Error; err != nil {
		return err
	}
	return nil
}


func GetSystemLogCount(c *gin.Context) (int, error) {
	count := 0
	DB, err := utils.InitDB()
	defer DB.Close()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("status = %d", 0)
	if err := DB.Debug().Model(&models.Log{}).Where(sql).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}