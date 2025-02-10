package gws_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lxzan/gws"
)

const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

func GwsTest() {
	upgrader := gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		ParallelEnabled:   true,                                 // 开启并行消息处理
		Recovery:          gws.Recovery,                         // 开启异常恢复
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // 开启压缩
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop() // 此处阻塞会使请求上下文不能顺利被GC
		}()
	})
	http.ListenAndServe(":8600", nil)
}

type Handler struct{}

// 连接事件
func (c *Handler) OnOpen(socket *gws.Conn) {
	fmt.Println(socket.RemoteAddr(), "连接成功")
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait)) // 设置该连接的心跳及断开
}

// 断开事件
func (c *Handler) OnClose(socket *gws.Conn, err error) {
	fmt.Println(socket.RemoteAddr(), "连接断开")
}

// 收到心跳
func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	fmt.Println("收到心跳", payload)
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait)) // 更新断开时间
	_ = socket.WritePong(nil)                                       // 返回心跳
}

// 返回心跳
func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {
	fmt.Println("返回心跳", payload)
}

// 收到信息
func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	fmt.Println("收到信息", message.Data)
	defer message.Close()
	socket.WriteMessage(message.Opcode, message.Bytes())
}
