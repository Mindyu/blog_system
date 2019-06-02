package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	upgrade *websocket.Upgrader
	connMap map[string]*WsConnection // 每一个用户对应一个chan
}

type WsConnection struct {
	msgChan chan string
}

var ws *WsServer

func init() {
	ws = new(WsServer)
	ws.upgrade = &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws.connMap = make(map[string]*WsConnection)
}

func GetWebSocket() *WsServer {
	return ws
}

func (self *WsServer) SendMsgToChan(user string, val string) {
	msc, exist := self.connMap[user]
	if exist {
		msc.msgChan <- val
	}
}

func (self *WsServer) HandleRequest(user string, w http.ResponseWriter, r *http.Request) {
	wsc := new(WsConnection)
	wsc.msgChan = make(chan string)
	self.connMap[user] = wsc

	wsc.ServeHTTP(w, r)

}

func (self *WsConnection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error:", err)
		return
	}
	fmt.Println("client connect :", conn.RemoteAddr())
	go self.connHandle(conn)

}

func (self *WsConnection) connHandle(conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	stopCh := make(chan int)
	// 服务器向浏览器发送消息
	self.send(conn, stopCh)
	// 接收消息
	go self.receive(conn, stopCh)
}

func (self *WsConnection) send(conn *websocket.Conn, stopCh chan int) {
	for {
		select {
		case <-stopCh:
			fmt.Println("connect closed")
			return
		case sender := <-self.msgChan:
			data := fmt.Sprintf("【%s】发来一条私信，请注意查收！", sender)
			err := conn.WriteMessage(1, []byte(data))
			fmt.Println("sending....")
			if err != nil {
				fmt.Println("send msg faild ", err)
				return
			}
		}
	}
}

func (self *WsConnection) receive(conn *websocket.Conn, stopCh chan int) {
	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Minute))
		_, msg, err := conn.ReadMessage() // 如果浏览器端断开，则无法读取数据，然后通过stopCh断开websocket连接
		if err != nil {
			close(stopCh)
			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					fmt.Printf("ReadMessage timeout remote: %v\n", conn.RemoteAddr())
					return
				}
			}
			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("ReadMessage other remote:%v error: %v \n", conn.RemoteAddr(), err)
			}
			return
		}
		fmt.Println("收到消息：", string(msg))
	}
}
