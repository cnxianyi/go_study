package fmt

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FmtTest(c *gin.Context) {
	// 基本打印
	fmt.Print("Hello")              // 打印
	fmt.Println("Hello")            // 打印并换行
	fmt.Printf("Hello %s", "World") // 格式化打印

	// 示例
	name := "Tom"
	age := 20
	fmt.Print("Name:", name)                    // Name:Tom
	fmt.Println("Age:", age)                    // Age: 20
	fmt.Printf("%s is %d years old", name, age) // Tom is 20 years old

	// 返回字符串
	s1 := fmt.Sprint("Hello")              // 转换为字符串
	s2 := fmt.Sprintln("Hello")            // 转换为字符串并添加换行
	s3 := fmt.Sprintf("Hello %s", "World") // 格式化字符串

	// 示例
	name = "Tom"
	s4 := fmt.Sprintf("Hello, %s!", name) // "Hello, Tom!"

	fmt.Println(s1, s2, s3, s4)

	// 扫描输入
	fmt.Scan(&name, &age)           // 空格分隔的输入
	fmt.Scanln(&name, &age)         // 一行输入
	fmt.Scanf("%s %d", &name, &age) // 格式化输入

	var value = 1

	// 1. 通用
	fmt.Printf("%v", value)  // 默认格式
	fmt.Printf("%+v", value) // 添加字段名
	fmt.Printf("%#v", value) // Go语法格式
	fmt.Printf("%T", value)  // 类型

	// 2. 布尔
	fmt.Printf("%t", true) // true 或 false

	// 3. 整数
	fmt.Printf("%d", 123) // 十进制
	fmt.Printf("%b", 123) // 二进制
	fmt.Printf("%o", 123) // 八进制
	fmt.Printf("%x", 123) // 十六进制（小写）
	fmt.Printf("%X", 123) // 十六进制（大写）

	// 4. 浮点数
	fmt.Printf("%f", 123.456)   // 默认精度
	fmt.Printf("%.2f", 123.456) // 保留2位小数
	fmt.Printf("%e", 123.456)   // 科学计数法
	fmt.Printf("%g", 123.456)   // 根据值选择 %e 或 %f

	// 5. 字符串
	fmt.Printf("%s", "hello") // 字符串
	fmt.Printf("%q", "hello") // 带引号字符串
	fmt.Printf("%x", "hello") // 十六进制

	// 6. 指针
	fmt.Printf("%p", &value) // 指针地址

	// 7. 宽度和精度
	fmt.Printf("%9d", 123)      // 宽度9，右对齐
	fmt.Printf("%-9d", 123)     // 宽度9，左对齐
	fmt.Printf("%09d", 123)     // 宽度9，零填充
	fmt.Printf("%.2f", 123.456) // 精度2

	c.JSON(200, "")
}
