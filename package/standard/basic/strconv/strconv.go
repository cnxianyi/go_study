package strconv_test

import (
	"fmt"
	"strconv"
)

func StrconvTest() {
	// string -> int
	i, _ := strconv.Atoi("123")                // string -> int
	fmt.Println(i)                             // 123
	i64, _ := strconv.ParseInt("123", 10, 64)  // string -> int64 基数10
	fmt.Println(i64)                           //123
	u64, _ := strconv.ParseUint("123", 10, 64) // string -> uint64
	fmt.Println(u64)                           // 123

	// string -> float
	f32, _ := strconv.ParseFloat("123.456", 32) // string -> float32
	fmt.Println(f32)                            // 123.45600128173828
	f64, _ := strconv.ParseFloat("123.456", 64) // string -> float64
	fmt.Println(f64)                            // 123.456

	// int -> string
	fmt.Println(strconv.Itoa(123))

	// int16|32 -> string
	s := strconv.Itoa(int(int16(123)))
	s1 := strconv.FormatInt(int64(int32(123)), 10)
	fmt.Println(s, s1)

	// float64|32 -> string
	f := strconv.FormatFloat(123.456, 'f', 2, 64)
	var fl32 float32 = 123.456
	f1 := strconv.FormatFloat(float64(fl32), 'f', 2, 32)
	fmt.Println(f, f1)

}
