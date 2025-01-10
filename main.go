package main

import (
	"go_study/config"
	"go_study/models"
	"go_study/package/pzap"
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

	pzap.Init() // zap 日志
	defer pzap.Logger.Sync()

	r.Run(":" + port)
}
