package proto_test

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

// 将proto转为go
// protoc --proto_path=protoc --go_out=out protoc/*.proto
// protoc --proto_path=. --go_out=package/third/proto package/third/proto/main.proto
/*
1. 将当前文件夹下的 main.proto 转为 go文件 protoc --proto_path=. --go_out=package/third/proto package/third/proto/main.proto
*/

func ProtoTest() {
	usr := User{
		Name: "ilya",
		Age:  18,
	}

	// 将res的值 序列化为 proto
	res, _ := proto.Marshal(&usr)
	fmt.Println(res) // [10 4 105 108 121 97 16 18]

	usrProto := []byte{10, 4, 105, 108, 121, 97, 16, 18}

	// 将proto为 反序列化
	var rsq User
	_ = proto.Unmarshal(usrProto, &rsq) // 将usrProto 反序列化后存入 rsq
	fmt.Println(rsq.Name)               // ilya

	// 模拟未知字段
	nUsr := NewUser{
		Name:  "new User",
		Age:   19,
		Other: 100,
	}
	nRes, _ := proto.Marshal(&nUsr)

	_ = proto.Unmarshal(nRes, &rsq)
	fmt.Println(rsq.Name) // new User
	fmt.Println(rsq.Age)  // 19

	// 24: 编码后的 int32类型字段3
	fmt.Println(rsq.unknownFields) // [24 100]
}
