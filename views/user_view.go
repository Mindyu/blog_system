package views

import (
	"github.com/Mindyu/blog_system/handlers"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/Mindyu/blog_system/utils/dbmeta"
	"github.com/Mindyu/blog_system/utils/md5"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(c *gin.Context){
	username := c.Request.PostFormValue("username")
	password := c.Request.PostFormValue("password")

	user, err := handlers.Login(c, username, password)
	if err!=nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	accessToken, err := handlers.TokenHelper(c, user)
	if err!=nil {
		utils.MakeErrResponse(c, "获取token失败")
		return
	}

	accessTokenObj := models.AccessToken{
		Access_token:accessToken,
		Token_type:"bearer",
	}

	utils.MakeOkResponse(c, accessTokenObj)
}


func GetUser(c *gin.Context) {
	id := c.Query("id")

	userId, err := strconv.Atoi(id)
	if err!=nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err!=nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, user)
}

func AddUser(c *gin.Context) {
	user := &models.User{}

	err := c.ShouldBindJSON(user)
	if err!=nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	user.Password = md5.EncryptPassword(user.Password)
	err = stores.SaveUser(c, user)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Query("id")

	userId, err := strconv.Atoi(id)
	if err!=nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err!=nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	updated := &models.User{}
	err = c.ShouldBindJSON(updated)
	if err!=nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	if err := dbmeta.Copy(user, updated); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	err = stores.SaveUser(c, user)
	if err != nil{
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	userId, err := strconv.Atoi(id)
	if err!=nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err!=nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	user.Status = 1

	if err := stores.SaveUser(c, user); err != nil{
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}
