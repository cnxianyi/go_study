package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(getCommand())
}

// 获取命令行参数
//
//	@description os.Args 是命令行切片. 执行 ./main a 则 os.Args为["./main" , "a"]
//	@return string
func getCommand() string {
	// var s string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += " " + os.Args[i]
	// }
	// return s

	// s1 := ""
	// for _, arg := range os.Args[1:] {
	// 	s1 += " " + arg
	// }
	// return s1

	return strings.Join(os.Args[1:], " ")
}
