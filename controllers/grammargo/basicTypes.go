package grammargo

/** 基础数据类型
1. 数字
int8
int16
int32
int64
int
uint

uintptr 没有指定具体大小,但足以容纳指针

float32
float64

复数
complex64
complex128
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i

% 取模  -5%3 和 -5%-3 都是 -2. 取决于-5的符号
/ 结果取决于是否全是整数 如 5/4 = 1

不同类型的数值进行计算,不被GO允许
应该将其转为同一类型 int(a)
	需注意大尺寸转换到小尺寸整数类型,可能会改变数值
	浮点数转换为整数 则会丢失所有小数部分

2. 字符串
len(s) 获取字符串长度
s[i] 获取对应位置的字符的 *字节值*
s[i , j] 截取中间的字符串 i 和 j 都支持忽略
字符串支持 + 相连

3. 布尔值
true | false
&& 和 || 支持短路 , 一旦确立布尔值,那么运算符右边的值不会再被计算
	var s string = ""
	println(s != "" && s[0] == 'x') // false
Go 的判断语句(if ...) 不会将非布尔值隐式转换
	i := 1
	if i {
		println(i) // non-boolean condition in if statement
	}

	可以是 if i != 0

*/

func basicTypes() {

}
