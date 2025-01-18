package grammarmongo

import (
	"context"
	"fmt"
	"go_study/database/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// FindAll 查询所有
func FindAll(c *gin.Context) {
	var mdb = mongodb.GetDB("account").Collection("account_info")
	var ctx = context.Background()
	var res []bson.M
	// 执行查询，获取游标. 用于遍历查询结果
	cursor, err := mdb.Find(ctx, bson.D{{}}) // {}:查询条件为空
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer cursor.Close(ctx) // 确保查询结束后关闭游标

	// 将查询结果解码到 res 中
	if err := cursor.All(ctx, &res); err != nil {
		fmt.Println("解析数据失败:", err)
		return
	}

	// 输出查询结果
	fmt.Println(res)

	c.JSON(http.StatusOK, res)
}

// 仅查询Id&account
func FindAllIdAccount(c *gin.Context) {
	// AccountInfo 定义返回的数据结构
	type accountInfo struct {
		ID      bson.ObjectID `bson:"_id"`
		Account string        `bson:"account"`
	}
	mdb := mongodb.GetDB("account").Collection("account_info")

	// 允许最长等待时间为 10 秒
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建查询选项
	opts := options.Find().SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "account", Value: 1},
	})

	// 执行查询
	cursor, err := mdb.Find(ctx, bson.D{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	defer cursor.Close(ctx)

	// 用结构体切片存储结果
	var accounts []accountInfo

	// 解码查询结果
	if err := cursor.All(ctx, &accounts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析数据失败"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// 查询指定id的account
func FindAccountById(c *gin.Context) {
	// 获取并验证 ID 参数
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID 参数不能为空"})
		return
	}

	// 转换字符串 ID 为 ObjectID
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 ID 格式"})
		return
	}

	// 定义返回的数据结构
	type AccountInfo struct {
		ID      bson.ObjectID `bson:"_id"`
		Account string        `bson:"account"`
	}

	mdb := mongodb.GetDB("account").Collection("account_info")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 筛选
	filter := bson.D{{Key: "_id", Value: objectID}}

	// 投影，返回的字段
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "account", Value: 1},
	}

	// 设置查询选项
	opts := options.Find().SetProjection(projection)

	// 查询
	cursor, err := mdb.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "查询失败",
			"detail": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)

	var accounts []AccountInfo

	// 解码
	if err := cursor.All(ctx, &accounts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "解析数据失败",
			"detail": err.Error(),
		})
		return
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到对应的账户"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
