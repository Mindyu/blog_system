package views

import (
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func QuerySystemAccessCount(c *gin.Context) {
	count, err := stores.GetAccessCount(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, count)
}

func GetSystemAccessWeek(c *gin.Context) {
	now := time.Now()
	time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	stats, err := stores.GetAccessStatisticsByWeek(c, now)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	/*if len(stats) == 7 {
		utils.MakeOkResponse(c, stats)
		return
	}*/
	statMap := map[string]int{}
	for _, stat := range stats {
		statMap[stat.Date] = stat.Count
	}
	// 对当天不存在记录的数据进行补充
	result := Result{}

	const base_format = "2006-01-02"
	d, _ := time.ParseDuration("-24h")
	var i time.Duration
	for i = 6; i >= 0; i-- {
		t := now.Add(i * d)
		str := t.Format(base_format)
		val, exist := statMap[str]
		if exist {
			result.Day += str[5:] + ","
			result.Count += strconv.Itoa(val) + ","
			continue
		}
		result.Day += str[5:] + ","
		result.Count += "0,"
	}
	result.Day = result.Day[:len(result.Day)-1]
	result.Count = result.Count[:len(result.Count)-1]

	utils.MakeOkResponse(c, result)
}

type Result struct {
	Day   string `json:"day"`
	Count string `json:"count"`
}
