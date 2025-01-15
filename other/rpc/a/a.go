package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
rpc: 远程过程调用

通常用于 分布式系统 微服务等之间的 通信方式

Go中 rpc方法 规定
	1. 必须是导出方法(函数首字母大写)
	2. 必须有两个参数. 第一个参数为输入参数. 第二个参数为输出参数. 输出参数是接收方的输出的指针
	3. 返回值是 error
*/

type HelloService struct{}

// @req 接收参数的指针地址.
func (p *HelloService) Hello(res string, req *string) error {
	// 修改指针对应的值
	*req = "hello" + res
	return nil
}

func main() {
	fmt.Println("rpc A is running 6666")

	// 注册一个 rpc 服务
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println("创建失败")
		return
	}

	// 使用 for循环 使服务端保持持续监听
	// 当监听到一个新连接后,将任务交与协程. 并重新开始监听
	for {
		// 创建一个TCP监听器.
		// listener是线程安全的. *不会出现连接被丢弃的情况 即上一个Accept()正在分发任务且下一个Accept()还未创建时. 新的rpc仍然可以被接收*
		conn, err := listener.Accept() // 在此阻塞等待连接
		if err != nil {
			fmt.Println("创建失败")
			return
		}

		// 进行处理客户端请求
		go rpc.ServeConn(conn)
	}

}
