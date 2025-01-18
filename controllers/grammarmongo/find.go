package grammarmongo

import (
	"context"
	"fmt"
	"go_study/database/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// FindAll 查询所有
func FindAll(c *gin.Context) {
	var mdb = mongodb.GetDB("go_study").Collection("account_info")
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
	mdb := mongodb.GetDB("go_study").Collection("account_info")

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

	mdb := mongodb.GetDB("go_study").Collection("account_info")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 筛选
	filter := bson.D{{Key: "_id", Value: objectID}}

	// 投影，返回的字段
	projection := bson.D{
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

	var accounts []accountInfo

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

	c.JSON(http.StatusOK, accounts[0].Account)
}

// 聚合查询
// 使用 account集合 _id 查询 account中的 accountId. 再聚合查询phone
func FindPhoneWithAccountId(c *gin.Context) {
	// 获取 accountId 参数
	accountId := c.Query("accountId")

	// 转换 accountId 为 ObjectID
	accountObjectId, err := bson.ObjectIDFromHex(accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid accountId format"})
		return
	}

	// 获取 MongoDB 客户端和集合
	mdb := mongodb.GetDB("go_study").Collection("account_info")
	ctx := context.Background()

	// 聚合查询：从 account 集合通过 _id 查找 accountId，再与 phone_info 集合进行连接
	// match: 匹配筛选对应字段. 相当于FindBy
	// unset: 忽略对应字段
	// unwind: 只显示
	// sort: 排序
	// limit:
	// lookup: 连接操作符
	pipeline := mongo.Pipeline{
		// 第一个阶段：从 account 集合获取对应 accountId
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: accountObjectId}}}}, // 查找 account_info中 _id = accountObjectId 的内容
		// 第二个阶段：使用 $lookup 来从 phone_info 集合查询对应的 phone 信息
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "phone_info"},        // 连接的 目标集合
			{Key: "localField", Value: "_id"},         // 匹配 当前集合的字段 _id
			{Key: "foreignField", Value: "accountId"}, // 匹配 外部集合中的字段 accountId
			// 即匹配 account_info 的_id == phone_info 中的 accountId
			{Key: "as", Value: "phoneDetails"}, // 聚合结果命名
		}}},
	}

	// 执行聚合查询 返回游标
	cursor, err := mdb.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("聚合查询失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "聚合查询失败"})
		return
	}
	defer cursor.Close(ctx)

	var result []bson.M
	// 读取所有文档到 result
	if err := cursor.All(ctx, &result); err != nil {
		fmt.Println("读取结果失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取结果失败"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result)
}
