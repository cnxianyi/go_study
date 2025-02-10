package buf

import (
	"bytes"
	"fmt"
)

/*
缓冲区
存储字节
*/

func Buf() {
	// var b bytes.Buffer
	// b = new(bytes.Buffer)
	b := bytes.NewBuffer([]byte{0x1, 0x2, 0x3})

	// 读取指定长度的 数据并返回. off向后偏移
	fmt.Println(b.Next(3))

	fmt.Println(b.Bytes())           // [] buf 为空
	fmt.Println(b.Len())             // 0
	fmt.Println(b.Available())       // 0
	fmt.Println(b.AvailableBuffer()) // []

	// 添加buf
	b.Write([]byte{0x1}) // 写入尾部
	b.WriteByte(0x2)     // 写入尾部
	b.WriteString("3")   // 写入尾部
	b.WriteRune(4)       // 写入一个 rune

	fmt.Println(b.Len())
	fmt.Println(b.Bytes())

	b1 := bytes.NewBuffer([]byte{0x5})
	b.WriteTo(b1) // 将b中数据 写入新 buf

	fmt.Println(b1.Bytes())

	b2 := [1]byte{}        // 创建一个长度的字节数组
	n, _ := b1.Read(b2[:]) // 将数组转为1长度的切片并写入
	fmt.Println(n)         // 1
	fmt.Println(b2)        // 5

	fmt.Println(b1.ReadByte())     // 读取一个长度 off向后偏移1
	fmt.Println(b1.ReadBytes('-')) // 读取第一个分隔符前的内容 off偏移
}
