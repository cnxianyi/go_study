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
|HSET key field value [field value ...]|新增和覆盖|HSET user:1001 name "Alice" age "25" city "New York"|
|HGET key field|获取指定字段的值|HGET user:1001 name|
|HGETALL key|获取所有字段和值|HGETALL user:1001|
|HKEYS key|获取所有字段名|HKEYS user:1001|
|HVALS key|获取所有字段值|HVALS user:1001|
|HEXISTS key field|检查字段是否存在|HEXISTS user:1001 name|
|HDEL key field [field ...]|删除字段|HDEL user:1001 age city|
|HLEN key|获取字段数量|HLEN user:1001|
|HINCRBY key field increment|增加值 整数|HINCRBY user:1001 age 5|
|HINCRBYFLOAT key field increment|增加值 浮点|HINCRBYFLOAT user:1001 balance 15.5|
|HMGET key field [field ...]|获取多个值|HMGET user:1001 name age|
|HMSET key field value [field value ...]|设置多个字段值||
*/

func HashTest(c *gin.Context) {

	var ctx = context.Background()
	var rdb = redis.GetDB()

	// HSet 返回新增的key. 更新不会显示
	res1 := rdb.HSet(ctx, "h1", "name", "ilya")
	fmt.Println(res1) // hset h1 name ilya: 1
	res1 = rdb.HSet(ctx, "h1", map[string]interface{}{
		"age":  19,
		"live": false,
	}) // redis中 true/false 会被存储为 1/0
	fmt.Println(res1) // hset h1 age 18 live true: 2
	res1 = rdb.HSet(ctx, "h1", []string{"name", "mike", "age", "20", "live", "1"})
	fmt.Println(res1) // hset h1 name mike age 20 live 1: 0

	// HGet 获取指定key的值
	res2 := rdb.HGet(ctx, "h1", "age")
	fmt.Println(res2) // hget h1 age: 20

	// HGETALL 获取所有 key和值
	res3 := rdb.HGetAll(ctx, "h1")
	fmt.Println(res3) // hgetall h1: map[age:20 live:1 name:mike]
	res4, _ := res3.Result()
	fmt.Println(res4["name"]) // mike

	// HKEYS 获取所有 key
	res5 := rdb.HKeys(ctx, "h1")
	fmt.Println(res5) // hkeys h1: [name age live]

	// HVALS 获取所有 值
	res6 := rdb.HVals(ctx, "h1")
	fmt.Println(res6) // hvals h1: [mike 20 1]

	// HEXISTS 检查key是否存在于该hash
	res7 := rdb.HExists(ctx, "h1", "name")
	fmt.Println(res7) // hexists h1 name: true

	// HDEL 删除该hash中的 key
	res8 := rdb.HDel(ctx, "h1", "age", "live")
	fmt.Println(res8) // hdel h1 age live: 2

	// HLEN 获取key 数量
	res9 := rdb.HLen(ctx, "h1")
	fmt.Println(res9) // hlen h1: 1

	// 增加 值 整数
	res10 := rdb.HIncrBy(ctx, "h1", "age", 1)
	fmt.Println(res10) // hincrby h1 age 1: 1

	// 增加 值 浮点数
	res11 := rdb.HIncrByFloat(ctx, "h1", "age", 1.1)
	fmt.Println(res11) // hincrbyfloat h1 age 1.1: 2.1

	// 获取多个字段值
	res12 := rdb.HMGet(ctx, "h1", "name", "age")
	fmt.Println(res12) // hmget h1 name age: [mike 2.1]

	// ~~设置多个字段~~ 不推荐. 应使用HSET
	res13 := rdb.HMSet(ctx, "h1", "name", "newname", "age", 18)
	fmt.Println(res13) // hmset h1 name newname age 18: true

	c.JSON(200, gin.H{
		"message": "success",
	})
}
