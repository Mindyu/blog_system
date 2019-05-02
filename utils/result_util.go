package utils

import (
	"github.com/Mindyu/blog_system/models/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeOkResponse(c *gin.Context, data interface{})  {
	c.JSON(http.StatusOK, common.Result{
		Status:"ok",
		Data:data,
	})
}

func MakeErrResponse(c *gin.Context, err string)  {
	c.JSON(http.StatusOK, common.Result{
		Status:"error",
		ErrMsg:err,
	})
}

func ValidErr(c *gin.Context, err error){
	if err!=nil {
		MakeErrResponse(c, err.Error())
	}
}

func ValidErrWithMsg(c *gin.Context, err error, msg string){
	if err!=nil {
		MakeErrResponse(c, msg)
	}
}