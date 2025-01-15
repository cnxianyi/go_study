package main

import (
	"go_study/config"
	"go_study/models"
	strconv_test "go_study/package/standard/basic/strconv"
	"go_study/package/third/zap"
	"go_study/practice/multithread"
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

	strconv_test.StrconvTest()

	println("_-----------------")

	multithread.PrintWithTwoThread()

	r.Run(":" + port)

}
