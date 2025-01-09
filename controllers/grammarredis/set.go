package grammarredis

import (
	"context"
	"fmt"
	"go_study/database/redis"

	"github.com/gin-gonic/gin"
)

/*
|命令|描述|示例|
|:---|:---|:---|
|SADD myset value [value ...]|添加元素|SADD myset "apple" "banana"|
|SMEMBERS key|获取所有|SMEMBERS myset|
|SISMEMBER key member|查询是否存在|SISMEMBER myset "apple"|
|SREM key member [member ...]|删除元素|SREM myset "apple"|
|SPOP key|随机移除一个元素|SPOP myset|
|SRANDMEMBER key [count]|随机获取|SRANDMEMBER myset 2|
*/

func SetTest(c *gin.Context) {

	var ctx = context.Background()
	var rdb = redis.GetDB()

	// 返回存入成功的数量
	res := rdb.SAdd(ctx, "myset", []string{"1", "2"})
	fmt.Println(res)                 // sadd myset 1 2: 2
	rdb.SAdd(ctx, "myset", "3", "4") // 支持连续参数

	// 获取所有
	res1 := rdb.SMembers(ctx, "myset")
	fmt.Println(res1) // smembers myset: [1 2 3 4]

	// 检查是否存在
	res2 := rdb.SIsMember(ctx, "myset", "3")
	fmt.Println(res2) // sismember myset 3: true

	// 删除元素
	res3 := rdb.SRem(ctx, "myset", 3)
	fmt.Println(res3) // srem myset 3: 1

	// 随机移除元素 返回被移除的元素
	res4 := rdb.SPop(ctx, "myset")
	fmt.Println(res4) // spop myset: 1

	// 随机获取一个元素 返回获取的元素
	res5 := rdb.SRandMember(ctx, "myset")
	fmt.Println(res5) // srandmember myset: 2

	// 获取set 长度
	res6 := rdb.SCard(ctx, "myset")
	fmt.Println(res6) // scard myset: 3

	// 返回多个集合的 交集
	rdb.SAdd(ctx, "myset1", "1", "2", "3", "4")
	res7 := rdb.SInter(ctx, "myset", "myset1")
	fmt.Println(res7) // sinter myset myset1: [1 2]

	// 并集
	res8 := rdb.SUnion(ctx, "myset", "myset1")
	fmt.Println(res8) // [1 2 3 4]

	// 差集 a中有 其他集合中没有的
	res9 := rdb.SDiff(ctx, "myset", "myset1")
	fmt.Println(res9) // sdiff myset myset1: []

	// 保存交集并集差集
	res10 := rdb.SInterStore(ctx, "myset1+1", "myset", "myset1")
	println(res10) // 0xc0005101e0 地址
	// ...

	// 删除集合
	res11 := rdb.Del(ctx, "myset", "myset1", "myset1+1")
	fmt.Println(res11) // del myset myset1 myset1+1: 3

	c.JSON(200, gin.H{
		"message": "success",
	})
}
