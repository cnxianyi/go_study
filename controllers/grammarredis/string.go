package grammarredis

import (
	"context"
	"fmt"
	"go_study/database/redis"

	"github.com/gin-gonic/gin"
)

/*
字符串
SET key value 设置键值
GET key 键获取值
DEL key 删除键值对
原子 SETNX key value 如果键不存在,则创建

// 下面的内容需要值为整数类型
INCR key 自增值
INCRBY key increment 指定步长
DECR key 自减
DECRBY key increment
//

APPEND 字符串追加. 返回字符串长度
STRLEN 字符串长度

// 位操作
SETBIT key offset value	设置 key 的指定偏移位为 0 或 1	SETBIT bitmap 7 1
GETBIT key offset	获取 key 的指定偏移位的值	GETBIT bitmap 7
BITCOUNT key	统计 key 中值为 1 的位数量	BITCOUNT bitmap
BITOP operation destkey key [key ...]

// 字符串操作
GETRANGE key start end 获取指定部分
SETRANGE key offset value 替换内容
*/
func StringTest(c *gin.Context) {
	var rdb = redis.GetDB()
	var ctx = context.Background()

	// SET key value
	// ctx , 键 , 值 , ttl
	rdb.Set(ctx, "stringTestSet", "value1", 0)

	// GET key
	res := rdb.Get(ctx, "stringTestSet")
	// res // get test1: value1 // res.args()... res.result()
	res1, _ := res.Result()
	fmt.Println(res1) // value1

	// DEL key 支持多个键同时删除
	res2 := rdb.Del(ctx, "string1TestSet") // 错误的键
	// fmt.Println(res2) //del stringTestSet: 1
	fmt.Println(res2.Result()) // 1 <nil> // 修改的行
	// 判断 第一个值 是否为0 即可
	if code, _ := res2.Result(); code == 0 {
		println("删除失败\n")
	}

	//SETNX key value
	res3 := rdb.SetNX(ctx, "stringTestSet", "value1", 0)
	// fmt.Println(res3) // setnx string1TestSet value1: true
	if boo, _ := res3.Result(); boo != true {
		println("键已经存在\n")
	}

	//INCR 将 key的值 自增. 要求值为整数类型
	rdb.Set(ctx, "incr1", "s", 0)
	res4 := rdb.Incr(ctx, "incr1")
	fmt.Println(res4) // incr incr1: ERR value is not an integer or out of range

	// INCRBY 指定步长自增
	rdb.Set(ctx, "incr1", "1", 0)
	res5 := rdb.IncrBy(ctx, "incr1", 2)
	fmt.Println(res5) // incrby incr1 2: 3

	// DECR 自减
	res6 := rdb.Decr(ctx, "incr1")
	fmt.Println(res6) // decr incr1: 2

	// DECRBY 指定步长自减
	res7 := rdb.DecrBy(ctx, "incr1", 2)
	fmt.Println(res7) // decrby incr1 2: 0

	rdb.Set(ctx, "stringTestSet1", "hello", 0)

	// APPEND 字符串拼接 返回字符串长度
	res8 := rdb.Append(ctx, "stringTestSet1", " world")
	fmt.Println(res8) // append stringTestSet1  world: 11

	// STRLEN 字符串长度
	res9 := rdb.StrLen(ctx, "stringTestSet1")
	fmt.Println(res9) // strlen stringTestSet1: 11

	// GETRANGE 获取部分字符串
	res10 := rdb.GetRange(ctx, "stringTestSet1", 0, 4)
	fmt.Println(res10) // getrange stringTestSet1 0 4: hello

	// SETRANGE 替换部分字符串
	res11 := rdb.SetRange(ctx, "stringTestSet1", 0, "HELLO")
	fmt.Println(res11) // setrange stringTestSet1 4 HELLO: 11

	res12, _ := rdb.Get(ctx, "stringTestSet1").Result()
	fmt.Println(res12) // HELLO world

	c.JSON(200, gin.H{
		"message": "success",
	})
}
