package grammarredis

import (
	"context"
	"fmt"
	"go_study/database/redis"

	"github.com/gin-gonic/gin"
)

/*
List 基于链表实现
可以实现简单的 消息队列
利用 List 的阻塞操作（如 BLPOP 和 BRPOP），可以实现任务分发和处理
使用 LPUSH 和 LPOP 模拟栈结构(先进后出)
使用 RPUSH 和 LPOP 模拟队列结构(先进先出)

// 插入
LPUSH key value [value ...] 插入头部
RPUSH key value [value ...] 插入尾部
LPUSHX key value 仅列表存在才会插入头部
RPUSHX key value 仅列表存在才会插入尾部

// 弹出
LPOP key 移除并返回 头部
RPOP key 移除并返回 尾部
RPOPLPUSH source dest 弹出尾部并插入到头部

// 获取
LRANGE key start stop 获取指定范围
LLEN key 获取列表长度
LINDEX key index 获取指定位置元素
LSET key index value 更新指定位置元素

// 删除
LREM key count value
LTRIM key start stop

// 阻塞
BLPOP key [key ...] timeout
BRPOP key [key ...] timeout
BRPOPLPUSH source dest timeout
*/
func ListTest(c *gin.Context) {

	var rdb = redis.GetDB()
	var ctx = context.Background()

	// 删除
	rdb.Del(ctx, "mylist")

	// Result 返回更新后的list长度
	res := rdb.LPush(ctx, "mylist", []string{"a", "b"}) // b a
	fmt.Println(res)                                    // lpush mylist a b: 2

	res = rdb.LPush(ctx, "mylist", "c") // c b a
	fmt.Println(res)                    // lpush mylist c: 3

	res = rdb.RPush(ctx, "mylist", "d") // c b a d
	fmt.Println(res)                    // rpush mylist d: 4

	// 仅列表存在才会插入成功
	res = rdb.LPushX(ctx, "mylist1", "e")
	fmt.Println(res) // lpushx mylist1 e: 0

	res = rdb.RPushX(ctx, "mylist", "e")
	fmt.Println(res) // rpushx mylist e: 5

	// LPOP 返回被弹出的内容
	res1 := rdb.LPop(ctx, "mylist")
	fmt.Println(res1) // lpop mylist: c

	// RPop 弹出尾部
	res1 = rdb.RPop(ctx, "mylist")
	fmt.Println(res1) // rpop mylist: e

	// RPopLPush 弹出mylist尾部元素 , 插入到newMylist 头部
	res1 = rdb.RPopLPush(ctx, "mylist", "newMylist")
	fmt.Println(res1) // rpoplpush mylist newMylist: d

	// LRange 获取指定位置的 -1 为最后一个
	res2 := rdb.LRange(ctx, "mylist", 0, -1)
	fmt.Println(res2)

	// LLEN 获取列表长度
	res3 := rdb.LLen(ctx, "mylist")
	fmt.Println(res3) // llen mylist: 2

	// LINDEX 指定位置元素
	res4 := rdb.LIndex(ctx, "mylist", 1)
	fmt.Println(res4) // lindex mylist 1: a

	// LSet 更新指定位置元素
	res5 := rdb.LSet(ctx, "mylist", 1, "f")
	cmd, err := res5.Result()
	if err != nil {
		fmt.Println(err) // ERR index out of range
	}
	fmt.Println(cmd) // OK
	fmt.Println(err) // <nil>
	//fmt.Println(res5) // lset mylist 1 f: OK

	// LRem 删除对应值的元素. 最多删除 几个
	// count为正整数则从头到尾 负数则从尾到头
	res6 := rdb.LRem(ctx, "mylist", 1, "f")
	fmt.Println(res6) // lrem mylist 1 f: 1

	// LTrim 保留指定范围
	res7 := rdb.LTrim(ctx, "mylist", 0, 1)
	fmt.Println(res7) // OK

	// 阻塞
	// 适合生产者-消费者模型
	// ...

	c.JSON(200, gin.H{
		"message": "success",
	})
}
