package grammarmongo

import (
	"context"
	"go_study/database/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 插入一个 account
func InsertOneAccount(c *gin.Context) {
	mdb := mongodb.GetDB("go_study").Collection("account_info")
	ctx := context.Background()

	act := c.Query("account")

	doc := accountInfo{Account: act}

	res, err := mdb.InsertOne(ctx, doc)
	if err != nil {
		c.JSON(400, res)
		return
	}

	c.JSON(200, res)
}

// 插入多个用户
func InsertManyAccount(c *gin.Context) {
	mdb := mongodb.GetDB("go_study").Collection("account_info")
	ctx := context.Background()

	a1 := c.Query("a1")
	a2 := c.Query("a2")

	doc := []accountInfo{{Account: a1}, {Account: a2}}

	res, err := mdb.InsertMany(ctx, doc)
	if err != nil {
		c.JSON(400, res)
		return
	}

	c.JSON(200, res)
}

// 插入新表
func InsertOnePhoneWithAccount(c *gin.Context) {
	mdb := mongodb.GetDB("go_study").Collection("phone_info")
	ctx := context.Background()

	accountId := c.Query("accountId")
	phone := c.Query("phone")

	newAccountId, err := bson.ObjectIDFromHex(accountId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid accountId format"})
		return
	}

	doc := PhoneInfo{AccountId: newAccountId, Phone: phone}

	res, err := mdb.InsertOne(ctx, doc)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, res)
}
