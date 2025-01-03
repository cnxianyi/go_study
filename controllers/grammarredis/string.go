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
	// ...
	c.JSON(200, gin.H{
		"message": "success",
	})
}
