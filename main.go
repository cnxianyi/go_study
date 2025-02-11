package main

import (
	"go_study/config"
	"go_study/database/mongodb"
	"go_study/database/mysql"
	"go_study/database/redis"
	"go_study/models"
	os_test "go_study/package/standard/basic/os"
	"go_study/package/third/zap"
	"go_study/practice/multithread"
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

	do()

	defer func() {
		mongodb.CloseDB()
		mysql.CloseDB()
		redis.CloseDB()
	}()

	r.Run(":" + port)
}

func do() {
	// multithreadTest() // 多线程
	// gws_test.GwsTest() // ws
	// buf.Buf() // buffer
	// proto_test.ProtoTest() // proto
	// binary_test.BinaryTest() // 字节序 小端|大端

	other()
}

// 多线程
func multithreadTest() {
	// 双线程输出 0~20
	// multithread.PrintWithTwoThread()

	// 多线程执行5个任务并捕获错误
	multithread.ConcurrentTestA()
}

// io
func ioTest() {
	os_test.OsTest()
}

func other() {
	// Viper测试
	// viper.ViperTest()

	// Strconv 测试
	// strconv_test.StrconvTest()

	// 反射
	// reflect_test.ReTest()
}
