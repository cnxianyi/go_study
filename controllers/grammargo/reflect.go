package grammargo

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

/*
reflect 反射
reflect.TypeOf 变量类型
reflect.Value 变量的值
reflect.Kind 变量具体的类型
*/
func ReflectTest(c *gin.Context) {

	var i32 int32 = 1
	// 获取变量类型
	t := reflect.TypeOf(i32)
	// 获取变量值
	v := reflect.ValueOf(i32)
	fmt.Println("Type:", t)  // int32
	fmt.Println("Value:", v) // str
	// 获取变量类型
	fmt.Println("Kind:", t.Kind()) // int32

	// 获取i32的指针
	v1 := reflect.ValueOf(&i32)
	// 判断 v1 的类型是否为指针
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem() // 获取指针指向的值
		v1.SetInt(2)   // 修改值
	}
	fmt.Println("Value:", i32)

	p1 := Persion{
		Name: "Ilya",
		Age:  18,
	}

	t1 := reflect.TypeOf(p1)
	fmt.Println("t1 Type:", t1) // grammargo.Persion

	// 遍历结构体
	for i := 0; i < t1.NumField(); i++ {
		field := t1.Field(i)
		fmt.Println(field.Name, ":", field.Type)
	}

	// 指针修改结构体值
	v2 := reflect.ValueOf(&p1).Elem()
	v2N := v2.FieldByName("Name")
	if v2N.IsValid() && v2N.CanSet() {
		v2N.SetString("Alice")
	}
	fmt.Println("p1Name:", p1.Name) // Alice

	// 获取结构体 标签值
	v3 := reflect.TypeOf(Persion{})
	field, _ := v3.FieldByName("Name")
	fmt.Println("Tag:", field.Tag.Get("json")) // name

	// 调用结构体方法
	v4 := reflect.ValueOf(&p1) // 传递指针类型
	method := v4.MethodByName("Man")
	if method.IsValid() {
		args := []reflect.Value{reflect.ValueOf(p1)} // 方法参数是 Persion 类型
		method.Call(args)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

type Persion struct {
	Name string `json:"name"`
	Age  int
}

func (p *Persion) Man(p1 Persion) {
	fmt.Println(p1.Name, "is", p1.Age)
}
