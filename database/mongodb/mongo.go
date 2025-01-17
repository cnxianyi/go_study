package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mdb *mongo.Client

func ConnectMongo() error {
	url := os.Getenv("MONGO")
	if url == "" {
		return fmt.Errorf("MONGO: 环境变量未设置")
	}

	// 设置 mongodb API版本. 关联服务器版本
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(url). // 设置要连接的url
						SetServerAPIOptions(serverAPI) // 设置连接的API
	var err error
	// 建立连接
	mdb, err = mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("MongoDB 连接错误%s", err.Error())
	}

	var result bson.M // MongoDB使用BSON格式

	if err := mdb.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return fmt.Errorf("MongoDB 连接错误%s", err.Error())
	}

	fmt.Println("MongoDB 连接成功")
	return nil
}

func GetDB() *mongo.Database {
	return mdb.Database("account")
}

func CloseDB() {
	if mdb != nil {
		if err := mdb.Disconnect(context.TODO()); err != nil {
			fmt.Printf("MongoDB 连接断开%s\n", err.Error())
		}
	}
}
