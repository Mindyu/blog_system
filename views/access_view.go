package views

import (
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func QuerySystemAccessCount(c *gin.Context) {
	count, err := stores.GetAccessCount(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, count)
}
