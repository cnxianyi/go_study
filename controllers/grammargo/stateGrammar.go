package grammargo

/* Go命名

保留字
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var

保留字
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64

	uint uint8 uint16 uint32 uint64 uintptr
	float32 float64 complex128 complex64
	bool byte rune string error

内建函数: make len cap new append copy close delete

	complex real imag
	panic recover

作用域:
	1. 在函数内定义,作用域在函数内
	2. 在函数外定义
		1. 首字母小写: 仅当前包
		2. 首字母大写: 会被包暴露
*/

/* Go 变量声明

1. 初始值
数值类型 0
布尔 false
字符串 ""
接口或引用类型 nil

2. 声明
var s string = "str"
var s, s1 = "s1", "s2"
s := "s1"

3. 指针
指针的值 是一个变量的 地址
&变量名 用于 获取该变量的地址(指针)
*变量名 用于 获取该地址的变量(值)

	TODO 交换指针
		x := 1
		y := 2
		// 交换 p q 的地址
		p := &x
		q := &y
		temp := &x
		p = q
		q = temp
		log.Println(*p) // 2
		log.Println(*q) // 1
		log.Println(x)  // 1
		log.Println(y)  // 2

		// 通过 指针改变对应变量的值
		*temp = 3
		log.Println(x) // 3

	TODO 指针之间的判断
		var x, y int
		// 默认值 与 本身
		println(x == x) // true
		println(x == y) // true

		// 指针之间 & 值相同但指针不同
		println(&x == &x)  // true
		println(&x == &y)  // true
		println(&x != nil) // true 指针不为空

	TODO 函数内的局部变量指针是安全的
		func c() *int {
			v := 1
			return &v
		}
		// 每次创建的指针都不同
		println(c(), c()) // 0xc00013f790 0xc00013f788

	TODO 指针 模拟C中的 increase
		// 1. 接收一个指针
		// 2. 将该指针对应的值自增
		// 3. 返回该指针对应的变量
		func addCount(p *int) int {
			*p++
			return *p
		}
		i := 0
		addCount(&i)
		addCount(&i)
		println(i) // 2

	TODO new
	new 语法糖会创建对应类型的默认值,且返回指针
		s := new(string)
		println(s) // 0xc0000bf788
		*s = "str"
		println(*s) // str
*/

func stateGrammar() {

}
