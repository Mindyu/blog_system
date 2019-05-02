package views

import (
	"github.com/Mindyu/blog_system/handlers"
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/Mindyu/blog_system/utils/md5"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
	"strings"
)

func Login(c *gin.Context) {
	username := c.Request.PostFormValue("username")
	password := c.Request.PostFormValue("password")

	user, err := handlers.Login(c, username, password)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	accessToken, err := handlers.TokenHelper(c, user)
	if err != nil {
		utils.MakeErrResponse(c, "获取token失败")
		return
	}

	accessTokenObj := common.AccessToken{
		Access_token: accessToken,
		Token_type:   "bearer",
	}

	utils.MakeOkResponse(c, accessTokenObj)
}

func QueryUserById(c *gin.Context) {
	id := c.Query("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, user)
}

func QueryAllUser(c *gin.Context) {
	param := &common.PageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	users, err := stores.GetUserList(c, param.CurrentPage, param.PageSize, param.RoleId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetUserListCount(c, param.RoleId, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum:total, List:users})
}

func QueryUserType(c *gin.Context) {
	roles, err := stores.GetAllRoles(c)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, roles)
}

func QueryUserAuth(c *gin.Context) {
	userName := c.Query("username")
	user, err := stores.GetUserByName(c, userName)
	if err != nil {
		utils.MakeErrResponse(c, "用户不存在")
	}
	roleId := user.RoleID
	//log.Info("user role:", roleId)
	roleAuthList, err := stores.GetRoleAuthByRoleID(c, roleId)
	if err != nil {
		utils.MakeErrResponse(c, "获取角色权限失败")
	}
	authIds := []int{}
	for _, roleAuth := range roleAuthList {
		authIds = append(authIds, roleAuth.AuthID)
	}
	authList, err := stores.GetAuthByIds(c, authIds)
	if err != nil {
		utils.MakeErrResponse(c, "获取权限失败")
	}
	utils.MakeOkResponse(c, authList)
}

func ValidUserName(c *gin.Context) {
	name := c.Param("name")

	user, _ := stores.GetUserByName(c, name)
	if user != nil {
		utils.MakeOkResponse(c, false)
		return
	} else {
		utils.MakeOkResponse(c, true)
		return
	}
}

func AddUser(c *gin.Context) {
	user := &models.User{}

	err := c.ShouldBindJSON(user)
	if err != nil {
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
	/*id := c.Query("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}*/

	updated := &models.User{}
	err := c.ShouldBindJSON(updated)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	/*if err := dbmeta.Copy(user, updated); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}*/
	idx := strings.Index(updated.Avatar, "?")
	updated.Avatar = updated.Avatar[:idx]

	err = stores.SaveUser(c, updated)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "更新成功")
}

func DeleteUserById(c *gin.Context) {
	id := c.Query("userID")

	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "用户ID转换失败")
		return
	}

	user, err := stores.GetUserByID(c, userId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	user.Status = 1

	if err := stores.SaveUser(c, user); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}
