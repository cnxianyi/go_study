package reflect_test

import (
	"fmt"
	"reflect"
)

/*
反射
*/
func ReTest() {
	r := Re{
		S: "s",
		I: 1,
	}

	// 反射 获取值
	fmt.Println("value:", reflect.ValueOf(r)) // {s 1}

	// typeof kind 反射 获取类型
	// typeof 返回的是用户类型 以及自定义数据. 如Tag
	// kind 返回的是底层类型
	fmt.Println("type:", reflect.TypeOf(r))        // reflect_test.Re
	fmt.Println("kind:", reflect.TypeOf(r).Kind()) // struct

	// valueOf 可以接收一个 可寻址的值. 只有这样,才能通过指针与反射修改 对应的值
	// 因此需要传入指针
	// 可寻址才能 操作该值. 否则只能读取
	// Elem 的前提就是 指针
	v := reflect.ValueOf(&r)
	fmt.Println("r value:", v) // &{s 1}
	// Elem() 用于获取反射出的 实际值
	v = v.Elem()
	fmt.Println("Elem r:", v) // {s 1}

	// FieldByName 根据字段名返回值
	fmt.Println("r.S:", v.FieldByName("S")) // s
	fmt.Println("r.I:", v.FieldByName("I")) // 1

	// FieldByIndex 根据index返回值
	fmt.Println("r i0:", v.FieldByIndex([]int{0})) // r i0: s
	fmt.Println("r i1:", v.FieldByIndex([]int{1})) // r i1: 1

	// isValid CanSet 检查是否存在以及 是否可以重写
	fmt.Println("isValid:", v.FieldByName("S").IsValid()) // true
	fmt.Println("CanSet:", v.FieldByName("S").CanSet())   // true

	// Set 修改值
	v.FieldByName("S").SetString("s1")
	v.FieldByName("I").SetInt(2)
	fmt.Println("r.S:", v.FieldByName("S")) // s1
	fmt.Println("r.I:", v.FieldByName("I")) // 2

	// Field 获取第几个字段
	fmt.Println("r. field 0:", v.Field(0)) // s1
	fmt.Println("r. field 1:", v.Field(1)) // 2

	// NumField 获取字段 Len
	f := v.NumField()
	fmt.Println("r NumField:", f) // 2

	// Go 中 Tag 存储在字段的元数据中
	// Type 获取反射的数据包含 StructField. Kind就不返回StructField. 必须是ValueOf返回的反射才能使用Type
	// Field 获取第几个字段的 StructField
	f1 := v.Type().Field(0)
	// 获取指定tag
	fmt.Println(f1.Tag.Get("tag")) // this is string

}

type Re struct {
	S string `tag:"this is string"`
	I int    `tag:"this is int"`
}
