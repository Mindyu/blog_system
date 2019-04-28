package utils

import (
	"github.com/Mindyu/blog_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeOkResponse(c *gin.Context, data interface{})  {
	c.JSON(http.StatusOK, models.Result{
		Status:"ok",
		Data:data,
	})
}

func MakeErrResponse(c *gin.Context, data interface{}, err string)  {
	c.JSON(http.StatusOK, models.Result{
		Status:"error",
		Data:data,
		ErrMsg:err,
	})
}