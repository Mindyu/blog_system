package views

import (
	"github.com/Mindyu/blog_system/middleware/jwt"
	"github.com/Mindyu/blog_system/server"
	"github.com/gin-gonic/gin"
	"time"
)

func WebSocket(c *gin.Context) {
	claims := c.MustGet("claims")
	customclaims := claims.(*jwt.CustomClaims)
	userName := customclaims.UserName

	ws := server.GetWebSocket()
	ws.HandleRequest(userName, c.Writer, c.Request)

	//go sendMsgTest()
}

func sendMsgTest() {
	for i := 0; i < 20; i++ {
		ws := server.GetWebSocket()
		ws.SendMsgToChan("mindyu", i)
		time.Sleep(time.Duration(time.Second * 5))
	}
}
