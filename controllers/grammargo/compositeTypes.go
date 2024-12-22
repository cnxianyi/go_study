package grammargo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

/*
数组的长度是 固定 的 , 在编译时就已经被固定
支持 下标访问 支持 len()
默认情况下 数组的元素会被初始化为 零值

	var arr [3]int = [3]int{1, 2}
	println(arr[2], len(arr)) // 0 3

可以使用 ... 来 自动创建值长度的数组

	arr := [...]int{1, 2, 3}
	println(len(arr)) // 3

定义部分位置的数组

	arr := []int{50: 50, 99: 1}
	println(len(arr), arr[50]) // 100 50
*/
func CompositeTypeByArray(c *gin.Context) {
	var arr1 [3]int = [3]int{1, 2}
	arr2 := [...]int{1, 2, 3}
	arr3 := []int{50: 50, 99: 1}

	c.JSON(http.StatusOK, gin.H{
		"arr1": arr1,
		"arr2": arr2,
		"arr3": arr3,
	})
}

/*
Slice是 对一个底层数组的 引用或视图

	Slice 由三个部分构成 指针、长度、容量
		指针: slice第一个元素对应的底层数组的地址
		长度: slice中元素的数目 *长度不能超过容量*
		容量: slice中开始位置 到 结束位置的位置

	Slice 的创建与数组类似,只是没有指定长度
		s := []int{1, 2}
		println(len(s), cap(s)) // 2 2

	Slice是底层数组的引用
		arr := [2]int{1, 2}
		s := arr[0:1]
		println(&s[0], &arr[0]) // 0xc0000bf788 0xc0000bf788 地址相同
		s[0] = 2
		println(s[0], arr[0]) // 2 2

	Slice 不支持 ==  Slice仅支持通过 == nil 来判断slice是否有底层数组
		arr := [2]int{1, 2}
		arr1 := [2]int{1, 2}
		println(arr == arr1) // true

		s1 := []int{1, 2}
		s2 := []int{1, 2}
		println(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)

		println(s1 == nil) // false
		s2 = nil
		println(s2 == nil) // true

		println(len([]int{})) // 0
	判断slice是否为零值应该使用 len来判断长度

	make 函数初始化slice 语法 make([]T , len)
		创建的默认为零值
		s := make([]int, 3)
		println(len(s), s[0]) // 3 0

	append 方法用于向slice 添加元素
		s := make([]int, 3)
		s = append(s, 1)
		println(len(s), s[3]) // 4 1
*/
func CompositeTypeBySlice(c *gin.Context) {
	s1 := []int{1, 2}
	s2 := make([]int, 3)
	s3 := make([]int, 3)
	s3 = append(s3, 1)
	println(len(s3), s3[3])

	c.JSON(http.StatusOK, gin.H{
		"s1": s1,
		"s2": s2,
		"s3": s3,
	})
}

/*
Map 哈希表

	无序的key/value 集合
	唯一key
	每个map都是 对一个哈希表的引用
	key 仅支持 支持 == 的类型作为key
	map可以通过测试key是否相等来判断是否已经存在
*/
func CompositeTypeByMap(c *gin.Context) {
	// make或字面量创建 Map
	m1 := make(map[string]int)
	m2 := map[string]int{}
	m1["m1String"] = 1
	println(m1)
	println(m2)
	// delete 删除
	m3 := map[string]int{"a": 1}
	delete(m3, "a")
	println(m3["a"]) // 0 零值

	// 无法 对map元素进行获取指针
	// map 元素的地址不能被取出，因为 map 的底层实现允许其内部数据结构在插入、删除或更新元素时动态扩展和重定位。也就是说，map 中的元素可能在操作时发生移动，因此直接取地址是不安全的。
	//m4 := map[string]int{"a": 1}
	// b1 := &m4["a"] // delcared and not used
	// println(b1) // invalid operation: cannot take address of m4["a"]

	// map 的遍历顺序是不确定的
	// map 的零值是 nil

	// 如何确定一个元素是否真的在map中,如这个值恰好是0
	m5 := map[string]int{"a": 0}
	if m, ok := m5["a"]; ok { // ok 的值是 map中是否真的有这个值
		println(m) // 0 只有真的存在才会执行,而不是返回零值
	}

	// map 之前也不能全等比较. 只能通过循环来实现
	//map 只能跟 nil ==

	m1["m1String"] = 1
	c.JSON(http.StatusOK, gin.H{
		"m1": m1,
		"m2": m2,
		"m3": m3,
		"m5": m5,
	})
}

/*
结构体

 1. 支持 . 访问

 2. 支持指针

 3. 支持 成员大写 自动导出

 4. 支持 结构体字面值 ilya := User{"Ilya", 18} | ilya := User{Name: "Ilya"} // {Ilya 0}

 5. 支持 函数返回值
    type User struct {
    Name string
    Age  int
    }

    func compositeType() User {
    ilya := User{Name: "Ilya"}
    return ilya // {Ilya 0}
    }

    将结构体作为参数传递到函数中时,Go中默认是值拷贝
    因此传递时应该使用指针

    type User struct {
    Name string
    Age  int
    }

    func test() *User {
    return &User{Name: "Ilya"}
    }

    func compositeType() {
    u1 := test()
    u1.Age = 18
    println(u1.Name, u1.Age) // Ilya 18
    }

    当结构体中 所有成员 都支持 比较,那么结构体之间也可以比较
    type User struct {
    Name string
    Age  int
    }
    u1 := User{Name: "U1", Age: 1}
    u2 := User{Name: "U2", Age: 1}
    println(u1 == u2)         // false
    println(u1.Age == u2.Age) // true
*/
func CompositeTypeByStruct(c *gin.Context) {
	type User struct {
		Name string
		Age  int
	}
	var ilya User = User{
		Name: "Ilya",
		Age:  18,
	}
	fmt.Println(ilya)                                         // {Ilya 18}
	println(ilya.Name + " is " + fmt.Sprintf("%d", ilya.Age)) // Ilya is 18
	ilya.Age = 19
	fmt.Println(ilya) // {Ilya 19}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"Ilya": ilya,
		},
	})
}

/*
JSON

  - json.Marshakl() 用于将结构体转换为json
  - json.Unmarshal() 用于将json转换为结构体
*/
func CompositeTypeByJson(c *gin.Context) {
	type User struct {
		Name string
	}
	j1 := `{"name": "ilya"}`
	var u1 User
	err := json.Unmarshal([]byte(j1), &u1)
	if err != nil {
		println("json解析失败")
	}
	println(u1.Name) // ilya
	j2, _ := json.MarshalIndent(u1, "", "  ")
	println(string(j2))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"u1.Name": u1.Name,
			"j2":      string(j2),
		},
	})
}

/*
模版字符串

  - 包 "text/template"
*/
func CompositeTypeByTextTemplate(c *gin.Context) {
	d1 := 1
	s := `s: {{.d1}}`

	tmpl, err := template.New("example").Parse(s)
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}

	data := map[string]interface{}{
		"d1": d1,
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		fmt.Println("Error Execute:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": result.String(),
	})
}
