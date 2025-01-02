package grammarredis

import (
	"context"
	"go_study/database/redis"
	"time"

	"github.com/gin-gonic/gin"
)

/*
go-redis

  - 默认 TTL 是纳秒为单位 建议使用time标准库
  - 所有方法都必须接受一个上下文参数
  - context.Context 标准库操作生命周期
  - context.Background() 默认创建空上下文
  - context.WithCancel(parentContext) 可以取消的上下文
  - context.WithTimeout 自动超时的上下文
*/
func RedisTest(c *gin.Context) {

	rdb := redis.GetDB()
	// 默认上下文. 可以直接使用
	rdb.Set(context.Background(), "test", "value", 10*time.Second)

	ctx := context.Background()

	// 设置两秒自动超时
	ctxWithTimeout, cancelTimeout := context.WithTimeout(ctx, 2*time.Second)

	defer cancelTimeout()

	// 设置超时
	err := rdb.Set(ctxWithTimeout, "test1", "value1", 10*time.Second)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to set test1",
			"error":   err.Err(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
