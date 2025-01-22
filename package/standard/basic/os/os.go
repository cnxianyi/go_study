package os_test

import (
	"fmt"
	"io"
	"os"
)

var testPath = "./package/standard/basic/os/test.txt"

func OsTest() {
	create()
	newFile()
	open()
	openFile()
	// osScan()
}

// 创建文件 默认权限0666
func create() {
	file, _ := os.Create(testPath)
	fmt.Printf("%v\n", file)
}

// 根据文件描述符创建对应文件 返回文件对象
func newFile() {
	os.NewFile(uintptr(1), testPath)
}

// 只读方式读取文件
func open() {
	file, err := os.Open(testPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, buf[:n]...)
	}

	fmt.Println(string(content))
}

// 配置权限以及打开方式 打开文件
func openFile() {
	file, err := os.OpenFile(testPath, 1, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
}

// 终端也是文件. 通过io进行获取 并打印
func osScan() {
	var buf [16]byte                     // 最多接收16个字符
	os.Stdin.Read(buf[:])                // 读取
	os.Stdin.WriteString(string(buf[:])) // 打印
}
