package grammarredis

import (
	"context"
	"fmt"
	"go_study/database/redis"
	"time"

	"github.com/gin-gonic/gin"
)

// 使用redis 实现一个简单的MQ 消息队列

// 消息队列处理通常是先进先出
// 可以使用 Redis中的 List 的RPUSH 和 LPOP 数据类型模拟

// 创建空接口
type Mq interface {
	// 初始化
	Init()

	// 消费者
	// 用于消费生产者发送的消息
	Consumer()

	// 生产者
	// 生产者用来产生消息并发送
	Producer()

	// 查询队列数据
	Ping()
}

type RedisMq struct{}

// 获取队列长度
func (RedisMq) Ping() string {
	var rdb = redis.GetDB()
	var ctx = context.Background()

	res, _ := rdb.LLen(ctx, "mq").Result()

	return fmt.Sprintf("队列剩余: %d", res)
}

// 添加队列. 当队列大于10时,模拟拒绝后续请求
func (RedisMq) Producer(s string) error {
	var rdb = redis.GetDB()
	var ctx = context.Background()

	len, _ := rdb.LLen(ctx, "mq").Result()

	if len > 10 {
		return fmt.Errorf("队列已满")
	}

	rdb.RPush(ctx, "mq", s)
	return nil
}

// 消费时应当 同一时间只能消费一个
func (RedisMq) Consumer() string {
	var rdb = redis.GetDB()
	var ctx = context.Background()

	// 使用 SetNX 模拟 竞用

	//获取锁
	lock := rdb.SetNX(ctx, "mq:lock", "l", 1)
	if boo, _ := lock.Result(); boo != true {
		return "消费失败,其他消费者正在消费"
	}

	res, _ := rdb.LPop(ctx, "mq").Result()

	if res == "" {
		return fmt.Sprintf("消费完成")
	}

	// 模拟1秒的锁
	rdb.Set(ctx, "mq:lock", "1", 1*time.Second)

	return fmt.Sprintf("消费了: %s", res)
}

func RedisMqPing(c *gin.Context) {
	var Mq RedisMq

	c.JSON(200, Mq.Ping())
}

func RedisMqProducer(c *gin.Context) {

	var Mq RedisMq

	type Produce struct {
		Name string `form:"name" binding:"required"`
	}

	var product Produce

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(400, "error")
		return
	}

	err = Mq.Producer(product.Name)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, "消费者添加成功")
}

func RedisMqConsumer(c *gin.Context) {
	var Mq RedisMq
	c.JSON(200, Mq.Consumer())
}
