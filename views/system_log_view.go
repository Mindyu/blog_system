package views

import (
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
)

func GetSystemLogList(c *gin.Context) {
	param := &common.LogPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	logs, err := stores.GetSystemLogList(c, *param)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetSystemLogListCount(c, *param)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: logs})
}


func GetSystemLogCount(c *gin.Context) {
	total, err := stores.GetSystemLogCount(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, total)
}

func DeleteSystemLogById(c *gin.Context) {
	id := c.Query("logId")

	logId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "ID转换失败")
		return
	}

	systemLog, err := stores.GetSystemLogByID(c, logId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	systemLog.Status = 1
	if err := stores.SaveSystemLog(c, systemLog); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

