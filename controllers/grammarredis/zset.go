package grammarredis

import (
	"context"
	"fmt"
	"go_study/database/redis"

	"github.com/gin-gonic/gin"
	redislib "github.com/redis/go-redis/v9" // 导入 redis 库
)

/*
|命令|描述|示例|
|:---|:---|:---|
|ZADD key [NX|XX] [CH] [INCR] score member [score member ...]|ZAdd 默认是 元素不存在则新增. 存在则更新score 返回新增的数量. 支持控件: NX: *仅添加 XX: 仅更新 CH: 新增或更新 INCR: 增加score*|ZADD myzset INCR 2 "apple"|
|ZRANGE key start stop [WITHSCORES]|获取所有元素 WITHSCORES: 按照升序|ZRANGE myzset 0 -1 WITHSCORES|
|ZREVRANGE key start stop [WITHSCORES]|获取所有元素 WITHSCORES: 按照降序|ZREVRANGE myzset 0 -1 WITHSCORES|
|ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]|按分值范围获取元素|ZRANGEBYSCORE myzset 1 2 WITHSCORES|
|ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]|按分值降序获取元素|ZREVRANGEBYSCORE myzset 3 1 WITHSCORES|
|ZRANK key member|获取元素排名 从0开始|ZRANK myzset "apple"|
|ZREVRANK key member|获取元素排名 从末尾开始(倒数)|ZREVRANK myzset "apple"|
|ZSCORE key member|获取元素分值|ZSCORE myzset "apple"|
|ZREM key member [member ...]|删除元素|ZREM myzset "banana"|
|ZREMRANGEBYRANK key start stop|删除指定范围的元素|ZREMRANGEBYRANK myzset 0 1|
|ZREMRANGEBYSCORE key min max|按分值范围删除|ZREMRANGEBYSCORE myzset 1 2|
|ZCARD key|获取 Zset 长度|ZCARD myzset|
|ZCOUNT key min max|获取分值范围内元素数量|ZCOUNT myzset 1 3|
|ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight ...] [AGGREGATE SUM|MIN|MAX]|计算并集并创建|ZUNIONSTORE result 2 zset1 zset2 WEIGHTS 1 2 AGGREGATE SUM|
|ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight ...] [AGGREGATE SUM|MIN|MAX]|计算交集并创建||
*/

func ZsetTest(c *gin.Context) {

	var ctx = context.Background()
	var rdb = redis.GetDB()

	// ZAdd 默认是 元素不存在则新增. 存在则更新score 返回新增的数量
	res1 := rdb.ZAdd(ctx, "myzset", redislib.Z{
		Score:  1,
		Member: "value1",
	})
	fmt.Println(res1) // zadd myzset 2 value1: 1

	// Zadd NX 仅添加新元素
	res2 := rdb.ZAddNX(ctx, "myzset", redislib.Z{
		Score:  1,
		Member: "value2",
	})
	fmt.Println(res2) // zadd myzset nx 1 value2: 1

	// Zadd XX 仅更新已有元素. 仅返回0
	res3 := rdb.ZAddXX(ctx, "myzset", redislib.Z{
		Score:  2,
		Member: "value3",
	})
	fmt.Println(res3) // zadd myzset xx 2 value2: 0

	// Zadd CH 返回受影响的元素数量
	res4 := rdb.ZAddArgs(ctx, "myzset", redislib.ZAddArgs{
		Ch: true,
		Members: []redislib.Z{
			{
				Score:  3,
				Member: "value3",
			},
		},
	})
	fmt.Println(res4) // zadd myzset ch 3 value3: 1

	// ZAddArgs 支持 NX XX LT GT Ch
	// NX: 只在元素不存在时添加
	// XX: 只在元素存在时更新
	// GT: 只在新分数大于当前分数时更新
	// LT: 只在新分数小于当前分时时更新
	res5 := rdb.ZAddArgs(ctx, "myzset", redislib.ZAddArgs{
		Ch: true,
		LT: true,
		Members: []redislib.Z{
			{
				Score:  2,
				Member: "value3",
			},
		},
	})
	fmt.Println(res5) // zadd myzset lt ch 2 value3: 1

	c.JSON(200, gin.H{
		"message": "success",
	})
}
