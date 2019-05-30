package stores

import (
	"fmt"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/persistence"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetAccessCount(c *gin.Context) (int, error) {
	count := 0

	if err := persistence.GetOrm().Debug().Model(&models.Access{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}


func SaveAccess(c *gin.Context, access *models.Access) error {

	if err := persistence.GetOrm().Save(access).Error; err != nil {
		return err
	}
	return nil
}

type Stat struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func GetAccessStatisticsByWeek(c *gin.Context, now time.Time) ([]*Stat, error) {
	stats := []*Stat{}

	sql := `SELECT 
	DATE_FORMAT(access_time,'%%Y-%%m-%%d') as date,count(*) as count
FROM access 
WHERE access_time > DATE_SUB('%s',INTERVAL 6 DAY)
GROUP BY date `
	sql = fmt.Sprintf(sql, now)
	sql = strings.Replace(sql, "\r\n", "\n", -1)
	if err := persistence.GetOrm().Debug().Raw(sql).Scan(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil

}