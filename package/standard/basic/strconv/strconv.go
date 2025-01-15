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
}
