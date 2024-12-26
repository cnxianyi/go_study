package grammargo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
方法
  - 一种特殊的函数,通常与struct关联.
  - 方法名在接收器（receiver）后面

方法值
  - 使用一个参数绑定具体的接收器
  - 调用时只需要调用本身

方法表达式
  - 绑定到类型本身,调用时需要传入接收器
*/
func MethodTest(c *gin.Context) {
	w := wife{
		Name: "Ilya",
		Age:  18,
	}
	w.life()        // 生活一年
	r1 := w.marry() // 方法
	r2 := marry(w)  // 普通函数

	// 方法值
	m1 := w.life
	m1()
	m1()
	r3 := w.marry()

	// 方法表达式
	m2 := (*wife).life
	r4 := &wife{
		Name: "Alice",
		Age:  18,
	}
	m2(r4)
	m2(r4)

	c.JSON(http.StatusOK, gin.H{
		"s1": r1, // 19
		"s2": r2, // 19
		"s3": r3, // 21
		"s4": r4, // 20
	})
}

// 创建一个 wife 的方法
func (w wife) marry() string {
	return fmt.Sprintf("恭喜你,你的结婚对象是 %d岁的 %s 女士!", w.Age, w.Name)
}

// 正常的marry函数
func marry(w wife) string {
	return fmt.Sprintf("恭喜你,你的虚空对象是 %d岁的 %s 女士!", w.Age, w.Name)
}

// 创建一个修改wife的方法 - 通过传入方法的指针
func (w *wife) life() {
	w.Age += 1
}

type wife struct {
	Name string
	Age  int
}

/*
扩展struct
*/
func ExpandStruct(c *gin.Context) {

	tom := cat{
		animal: animal{
			eat: "鱼",
		},
		bray: "喵",
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "扩展struct",
		"eat":     tom.eat,
		"bray":    tom.bray,
	})
}

type animal struct {
	eat string
}

type cat struct {
	animal
	bray string
}
