package views

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/persistence/trie"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
)

func GetBlogList(c *gin.Context) {
	param := &common.BlogPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)
	// 将JWT中的用户信息取出，如果为非管理员，则只能浏览和自己相关的博客
	param.Author = utils.InjectUserName(c)

	_, exist := c.Get("claims")
	if exist {
		blogs, err := stores.GetBlogList(c, param.CurrentPage, param.PageSize, param.BlogTypeId, param.SearchWords, param.Author, param.SortType, 0)
		if err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}
		total, err := stores.GetBlogListCount(c, param.BlogTypeId, param.SearchWords, param.Author, 0)
		if err != nil {
			utils.MakeErrResponse(c, err.Error())
			return
		}

		utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: blogs})
		return
	}
	// 游客访问，只能访问公有博客
	blogs, err := stores.GetBlogList(c, param.CurrentPage, param.PageSize, param.BlogTypeId, param.SearchWords, param.Author, param.SortType, 1)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetBlogListCount(c, param.BlogTypeId, param.SearchWords, param.Author, 1)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: blogs})
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

	// 阅读量+1
	blog.ReadCount += 1
	_ = stores.SaveBlog(c, blog)

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

func QueryBlogTags(c *gin.Context) {

	tags, err := stores.GetBlogTags(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	tagMap := map[string]int{}
	for _, tag := range tags {
		tagArr := strings.Split(tag.Keywords, ",")
		for _, one := range tagArr {
			val, exist := tagMap[one]
			if !exist {
				tagMap[one] = 1
				continue
			}
			tagMap[one] = val + 1
		}
	}

	tagList := []*common.Tag{}
	for key, val := range tagMap {
		tagList = append(tagList, &common.Tag{TagName: key, Count: val})
	}
	utils.MakeOkResponse(c, tagList)
}

func GetBlogSearchKeySug(c *gin.Context) {
	searchKey := c.Query("key")
	res := trie.GetKeyTrie().GetStartsWith(searchKey)
	utils.MakeOkResponse(c, res)
}
