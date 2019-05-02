package views

import (
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func QueryAllBlogType(c *gin.Context) {
	types, err := stores.GetAllBlogType(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, types)
}
