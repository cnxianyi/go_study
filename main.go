package main

import (
	"go_study/config"
	"go_study/database/mongodb"
	"go_study/database/mysql"
	"go_study/database/redis"
	"go_study/models"
	"go_study/package/third/zap"
	"go_study/router"
	"os"
)

func main() {
	config.SetupEnv()

	models.InitMysql()
	models.InitRedis()
	models.InitMongo()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认值8080
	}

	r := router.SetupRouter()

	zap.Init() // zap 日志
	defer zap.Logger.Sync()

	// Strconv 测试
	// strconv_test.StrconvTest()

	// 双线程输出 0~20
	// multithread.PrintWithTwoThread()

	// Viper测试
	// viper.ViperTest()

	// os_test.OsTest()

	defer func() {
		mongodb.CloseDB()
		mysql.CloseDB()
		redis.CloseDB()
	}()

	r.Run(":" + port)
}
