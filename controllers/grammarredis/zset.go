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

	// 不支持 WITHSCORES. 需要使用args
	res6 := rdb.ZRange(ctx, "myzset", 0, 2)
	fmt.Println(res6) // zrange myzset 0 2: [value1 value2 value3]

	// 设置 WITHSCORES
	res7 := rdb.ZRangeWithScores(ctx, "myzset", 0, 2)
	fmt.Println(res7) // zrange myzset 0 2 withscores: [{1 value1} {1 value2} {2 value3}]

	// 降序 查询 Range
	res8 := rdb.ZRevRange(ctx, "myzset", 0, 2)
	fmt.Println(res8) // zrevrange myzset 0 2: [value3 value2 value1]

	res9 := rdb.ZRevRangeWithScores(ctx, "myzset", 0, 2)
	fmt.Println(res9) // zrevrange myzset 0 2 withscores: [{2 value3} {1 value2} {1 value1}]

	// ZRANGEBYSCORE 按照分值范围获取
	res10 := rdb.ZRangeByScore(ctx, "myzset", &redislib.ZRangeBy{
		Max: "1",
		Min: "0",
	})
	fmt.Println(res10) // zrangebyscore myzset 0 1: [value1 value2]

	// ZREVRANGEBYSCORE 分值范围降序
	res11 := rdb.ZRevRangeByScore(ctx, "myzset", &redislib.ZRangeBy{
		Max: "1",
		Min: "0",
	})
	fmt.Println(res11) // zrevrangebyscore myzset 1 0: [value2 value1]

	// 支持withscores
	res12 := rdb.ZRangeByScoreWithScores(ctx, "myzset", &redislib.ZRangeBy{
		Max: "1",
		Min: "0",
	})
	fmt.Println(res12) // zrangebyscore myzset 0 1 withscores: [{1 value1} {1 value2}]

	res13 := rdb.ZRevRangeByScoreWithScores(ctx, "myzset", &redislib.ZRangeBy{
		Max: "1",
		Min: "0",
	})
	fmt.Println(res13) // zrevrangebyscore myzset 1 0 withscores: [{1 value2} {1 value1}]

	// ZRank 查询值排名
	res14 := rdb.ZRank(ctx, "myzset", "value3")
	fmt.Println(res14) // zrank myzset value2: 2

	// ZSCORE 获取指定元素的值
	res15 := rdb.ZScore(ctx, "myzset", "value3")
	fmt.Println(res15) // zscore myzset value3: 2

	// ZRem 删除元素 返回删除的数量
	res16 := rdb.ZRem(ctx, "myzset", "value1", "value2")
	fmt.Println(res16) // zrem myzset value1 value2: 2

	// ZREMRANGEBYRANK 删除指定 排名 的元素
	res17 := rdb.ZRemRangeByRank(ctx, "myzset", 0, 1)
	fmt.Println(res17) // zremrangebyrank myzset 0 1: 1

	// ZREMRANGEBYSCORE 删除指定 范围 的元素
	res18 := rdb.ZRemRangeByScore(ctx, "myzset", "0", "1")
	fmt.Println(res18) // zremrangebyscore myzset 0 1: 0

	rdb.ZAdd(ctx, "myzset", redislib.Z{
		Score:  1,
		Member: "value1",
	}, redislib.Z{
		Score:  2,
		Member: "value2",
	})

	// 获取长度
	res19 := rdb.ZCard(ctx, "myzset")
	fmt.Println(res19) // zcard myzset: 2

	// 获取值指定范围 的元素的数量
	res20 := rdb.ZCount(ctx, "myzset", "0", "1")
	fmt.Println(res20) // zcount myzset 0 1: 1

	rdb.ZAdd(ctx, "myzset1", redislib.Z{
		Score:  1,
		Member: "value1",
	}, redislib.Z{
		Score:  3,
		Member: "value3",
	})

	// 1. 基本的 ZUNIONSTORE
	res21, err := rdb.ZUnionStore(ctx, "destination", &redislib.ZStore{
		Keys: []string{"myzset1", "myzset"},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Union members count: %d\n", res21)

	// 2. 使用权重的 ZUNIONSTORE
	res22, err := rdb.ZUnionStore(ctx, "destination_weighted", &redislib.ZStore{
		Keys:    []string{"myzset1", "myzset"},
		Weights: []float64{2, 1}, // myzset1 的分数权重为 2，myzset 的分数权重为 1
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Weighted union members count: %d\n", res22)

	// 3. 使用聚合函数的 ZUNIONSTORE
	res23, err := rdb.ZUnionStore(ctx, "destination_min", &redislib.ZStore{
		Keys:      []string{"myzset1", "myzset"},
		Aggregate: "MIN", // 可以是 "SUM"(默认), "MIN", "MAX"
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Min aggregate union members count: %d\n", res23)

	// ZINTERSTORE 交集

	// 1. 基本的 ZINTERSTORE
	res24, err := rdb.ZInterStore(ctx, "inter_result", &redislib.ZStore{
		Keys: []string{"myzset1", "myzset"},
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Intersection count: %d\n", res24)

	// 2. 带权重的 ZINTERSTORE
	res25, err := rdb.ZInterStore(ctx, "inter_weighted", &redislib.ZStore{
		Keys:    []string{"myzset1", "myzset"},
		Weights: []float64{2, 1}, // myzset1 的分数权重为 2，myzset2 的分数权重为 1
	}).Result()
	fmt.Printf("Intersection count: %d\n", res25)

	// 3. 使用不同聚合函数
	res26, err := rdb.ZInterStore(ctx, "inter_min", &redislib.ZStore{
		Keys:      []string{"myzset1", "myzset"},
		Aggregate: "MIN", // 可以是 "SUM"(默认), "MIN", "MAX"
	}).Result()
	fmt.Printf("Intersection count: %d\n", res26)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
