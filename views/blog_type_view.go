package views

import (
	"github.com/Mindyu/blog_system/cache"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
)

func QueryAllBlogType(c *gin.Context) {
	/*types, err := stores.GetAllBlogType(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}*/
	utils.MakeOkResponse(c, cache.Cache())
}

func QueryBlogTypeStats(c *gin.Context) {
	types, err := stores.GetBlogTypeStats(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, types)
}

func AddBlogType(c *gin.Context) {
	blogType := &models.BlogType{}

	err := c.ShouldBindJSON(blogType)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	err = stores.SaveBlogType(c, blogType)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	// 新增类别之后，插入缓存
	cache.UpdateCache(blogType)

	utils.MakeOkResponse(c, blogType)
}