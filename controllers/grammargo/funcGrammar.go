package grammargo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
函数

函数声明: func 函数名 (参数列表) [(返回值列表)] {}
  - 如果设置了返回值,那么函数必须以return结尾
  - 函数的参数不支持默认值、且参数的传递必须严格按照声明顺序
  - 参数传递的是拷贝,但如果传递指针的话,函数是能够影响外部的
  - 支持可变参数args (args ...int)

defer
  - defer的调用时机是 当前函数逻辑结束,然后开始执行defer语句.
  - 且遵循先进后出 后进先出 LIFO的顺序
  - 通常用于执行 关闭文件,释放锁 类似finally?

panic异常
  - panic发生时,程序会中断运行,引起程序崩溃
  - panic 支持手动调用
  - 通常用于表述 程序到达了错误的路径

recover捕获异常
  - recover 必须在一个 defer函数中调用
  - 能够捕获到 panic的信息并处理
*/
func FuncGrammar(c *gin.Context) {
	deferTest()
	defer panicTest()
	recoverPanic()
	res := doubleJump()()
	args1 := argFunc(1, 2, 3, 4, 5)
	arr1 := []int{1, 2, 3, 4}
	args2 := argFunc(arr1...)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"res":   res,
			"args1": args1,
			"args2": args2,
		},
	})
}

func doubleJump() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func argFunc(args ...int) int {
	total := 0
	for _, i := range args {
		total += i
	}

	return total
}

func deferTest() {
	println(1)
	defer println(2)
	defer println(3)
	println(4)
	// 1 4 3 2
}

func panicTest() {
	panic("this is panic test")
}

func recoverPanic() {
	println("-----------")
	println(1)
	defer func() {
		if r := recover(); r != nil { // 捕获 panic 的信息
			fmt.Printf("%v\n", r) // 打印捕获到的 panic 信息
			println(2)
		}
	}()
	println(3)
	panic("some panic") // 触发 panic，跳转到 defer 执行
	println(4)
	// 1 3 some panic 2
}
