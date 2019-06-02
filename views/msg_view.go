package views

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/server"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func QueryNotReadMsgByName(c *gin.Context) {
	name := c.Query("name")

	msg, err := stores.GetNotReadMsgByUserName(c, name)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, msg)
}

func QueryReadMsgByName(c *gin.Context) {
	name := c.Query("name")

	msg, err := stores.GetReadMsgByUserName(c, name)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, msg)
}

func AddPrivateMsg(c *gin.Context) {
	msg := &models.PrivateMsg{}

	err := c.ShouldBindJSON(msg)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	msg.CreatedAt = time.Now()
	msg.Status = 0
	msg.IsRead = 0      // 设置未读
	err = stores.SavePrivateMsg(c, msg)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	ws := server.GetWebSocket()
	ws.SendMsgToChan(msg.Receiver, msg.Sender)

	utils.MakeOkResponse(c, msg)
}

func ReadPrivateMsg(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err!=nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	msg, err := stores.GetPrivateMsgById(c, id)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	msg.IsRead = 1
	err = stores.SavePrivateMsg(c, msg)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, msg)
}