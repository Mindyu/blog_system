package views

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
)

func GetBlogList(c *gin.Context){
	param := &common.BlogPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	blogs, err := stores.GetBlogList(c, param.CurrentPage, param.PageSize, param.BlogTypeId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetBlogListCount(c, param.BlogTypeId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum:total, List:blogs})

}


func QueryBlogById(c *gin.Context) {
	id := c.Query("blogId")

	blogId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "博客ID转换失败")
		return
	}

	blog, err := stores.GetBlogById(c, blogId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, blog)
}


func AddBlog(c *gin.Context) {
	blog := &models.Blog{}

	err := c.ShouldBindJSON(blog)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	err = stores.SaveBlog(c, blog)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, blog)
}

func UpdateBlog(c *gin.Context) {
	updated := &models.Blog{}
	err := c.ShouldBindJSON(updated)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	err = stores.SaveBlog(c, updated)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "更新成功")
}

func DeleteBlogById(c *gin.Context) {
	id := c.Query("blogId")

	blogId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	blog, err := stores.GetBlogById(c, blogId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	blog.Status = 1

	if err := stores.SaveBlog(c, blog); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func QueryBlogByMonth(c *gin.Context) {

	stats, err := stores.GetBlogStatsByMonth(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, stats)
}