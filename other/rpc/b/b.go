package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 创建一个 rpc 连接
	client, err := rpc.Dial("tcp", ":6666")
	if err != nil {
		fmt.Println("连接失败")
	}

	var req string
	// 第一个参数: 指定服务名和方法
	// 第二个参数: 调用参数
	// 第三个参数: 用于接收返回的结果
	err = client.Call("HelloService.Hello", " this is B", &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req) // hello this is B

}
