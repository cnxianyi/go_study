package main

import (
	"go_study/config"
	"go_study/models"
	"go_study/package/third/zap"
	"go_study/router"
	"os"
)

func main() {
	config.SetupEnv()

	models.InitMysql()
	models.InitRedis()

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

	r.Run(":" + port)

}
