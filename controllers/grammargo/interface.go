package grammargo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
接口 interface
 1. 接口是一种抽象的类型. 描述了一系列方法的合集
 2. 任何类型只要实现了接口中定义的所有方法.就被视为实现了该接口
 3. 接口支持设置为空 - 空接口不含renew方法.
    因为所有类型都隐式实现了空接口.
    所以空接口可以表示任何值.
 4. 可以使用类型断言判断接口的具体类型
 5. 接口支持嵌套
 6. 接口的零值是 nil
 7. 接口会保存 动态类型 和 动态值
*/
func InterfaceTest(c *gin.Context) {

	// 此时 I1类型已经实现了helloWorld的所有方法
	// I1 实现了 helloWorld 接口
	var h1 helloWorld = I1{}
	h1.hello()
	h1.world()
	res := h1.live("Ilya", 18)
	println(res)
	// 动态类型
	// h1 的动态类型:grammargo.I1 , 动态值: {}
	fmt.Printf("h1 的动态类型:%T , 动态值: %v\n", h1, h1)

	// 空接口
	var emptyInterface I2
	emptyInterface = 1
	fmt.Println(emptyInterface) // 1
	emptyInterface = false
	fmt.Println(emptyInterface) // false

	// 断言空接口的类型
	if _, ok := emptyInterface.(bool); ok {
		fmt.Println("是bool类型") // ✅
	} else {
		fmt.Println("是其他类型")
	}

	// 嵌套接口
	var h3 I3 = I1{}
	h3.hello() // hello
	// 动态类型
	// h3 的动态类型:grammargo.I1 , 动态值: {}
	fmt.Printf("h3 的动态类型:%T , 动态值:%v\n", h3, h3)
	// 回收接口
	var h4 I3 = I1{}
	h4 = nil
	fmt.Println(h4) // <nil>

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}

// 创建helloWorld接口. 其中需要实现两个方法
type helloWorld interface {
	hello()
	world()
	live(string, int) string
}

// 创建 I1 结构
type I1 struct{}

// I1结构实现 部分方法
func (I1) hello() {
	println("hello")
}

// I1结构实现 部分方法
func (I1) world() {
	println("world")
}

// I1结构 设置传入参数类型和返回值类型
func (I1) live(name string, live int) string {
	return fmt.Sprintf("%s lived for %d years", name, live)
}

// 创建一个空接口
type I2 interface{}

// 创建接口 嵌套其他接口
type I3 interface {
	helloWorld
}
