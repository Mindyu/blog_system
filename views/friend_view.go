package views

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/Mindyu/blog_system/models/common"
	"github.com/Mindyu/blog_system/stores"
	"github.com/Mindyu/blog_system/utils"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
)

func GetFriendList(c *gin.Context) {
	param := &common.RelationPageRequest{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}
	log.Info(param)

	friends, err := stores.GetFriendList(c, param.CurrentPage, param.PageSize, param.Username)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	total, err := stores.GetFriendListCount(c, param.Username)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	friendNames := []string{}
	for _, friend := range friends {
		friendName := friend.Username1
		if friendName == param.Username {
			friendName = friend.Username2
		}
		friendNames = append(friendNames, friendName)
	}
	user, err := stores.GetUsersByNamesAndLikeWord(c, friendNames, param.SearchWords)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, common.PageResult{TotalNum: total, List: user})
}

func DeleteFriendById(c *gin.Context) {
	id := c.Query("id")

	friendId, err := strconv.Atoi(id)
	if err != nil {
		utils.MakeErrResponse(c, "ID转换失败")
		return
	}

	friend, err := stores.GetFriendByID(c, friendId)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	friend.Status = 1
	if err := stores.SaveFriend(c, friend); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func DeleteFriendByName(c *gin.Context) {
	req := &struct {
		UserName   string `json:"user_name"`
		FriendName string `json:"friend_name"`
	}{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Info(err.Error())
		utils.MakeErrResponse(c, "参数解析失败")
		return
	}

	friend, err := stores.GetFriendByName(c, req.UserName, req.FriendName)
	if err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	friend.Status = 1
	if err := stores.SaveFriend(c, friend); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}
	utils.MakeOkResponse(c, "删除成功")
}

func AddFriend(c *gin.Context) {
	friend := &models.Friend{}

	err := c.ShouldBindJSON(friend)
	if err != nil {
		utils.MakeErrResponse(c, "好友信息转换失败")
		return
	}

	t, _ := stores.GetFriendByName(c, friend.Username1, friend.Username2)
	if t != nil {
		utils.MakeErrResponse(c, "已成为好友，请勿重复添加")
		return
	}

	friend.Status = 2      // 待另一个用户同意添加好友
	friend.CreatedAt = time.Now()

	if err := stores.SaveFriend(c, friend); err != nil {
		utils.MakeErrResponse(c, err.Error())
		return
	}

	utils.MakeOkResponse(c, "添加好友成功")
}
