package views

import (
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
)

func GetAttentionList(c *gin.Context) {
	param := &common.PageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	attentions, err := stores.GetAttentionList(c, param.CurrentPage, param.PageSize, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetAttentionListCount(c, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	attentionNames := []string{}
	for _, attention := range attentions {
		attentionNames = append(attentionNames, attention.FocusedUser)
	}
	user, err := stores.GetUsersByNames(c, attentionNames)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: user})
}

func DeleteAttentionById(c *gin.Context) {
	id := c.Query("id")

	attentionId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "ID转换失败")
		return
	}

	attention, err := stores.GetAttentionByID(c, attentionId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	attention.Status = 1
	if err := stores.SaveAttention(c, attention); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func DeleteAttentionByName(c *gin.Context) {
	req := &struct {
		UserName      string `json:"user_name"`
		AttentionName string `json:"attention_name"`
	}{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	attention, err := stores.GetAttentionByName(c, req.UserName, req.AttentionName)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	attention.Status = 1
	if err := stores.SaveAttention(c, attention); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}
