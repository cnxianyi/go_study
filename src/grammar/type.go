package grammar

import "fmt"

func AllType() {

	// var 变量 | const 常量

	// bool
	const b1 bool = true

	// int(32 or 64) 根据平台计算，32位则为-2,147,483,648 到 2,147,483,647
	// 64位则-9,223,372,036,854,775,808 到 9,223,372,036,854,775,807
	const i32_64 int = 1

	// int8, int16, int32, int64
	const i8 int8 = 127                   // -128 到 127
	const i16 int16 = 32767               // -32,768 到 32,767
	const i32 int32 = 2147483647          // -2,147,483,648 到 2,147,483,647
	const i64 int64 = 9223372036854775807 // -9,223,372,036,854,775,808 到 9,223,372,036,854,775,807

	// uint(32 or 64), uint8(byte), uint16, uint32, uint64
	// 无符号整数
	// 大小等于对应int的正整数部分乘二再加一 (127*2)+1 = 0~255
	const uint1 uint = 1

	// float32, float64
	const f32 float32 = 3.40e+38  // 约 ±1.18e-38 到 ±3.40e+38
	const f64 float64 = 1.80e+307 // 默认浮点类型 ±2.23e-308 到 ±1.80e+307

	// byte: uint8别名
	const byte1 byte = 255

	// rune: int32 别名
	const rune1 rune = 2147483647

	// string
	// complex64, complex128
	// array

	fmt.Println(byte1)

}
