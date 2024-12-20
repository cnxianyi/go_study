package main

import (
	"fmt"
	"go_study/router"
	dbMysql "go_study/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 配置env
func setupEnv() {
	// godotenv
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
		log.Print("当前环境", env)
	} else {
		log.Print("当前环境", env)
	}

	// 加载默认文件 .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(".env 文件未找到")
	}

	// 加载环境文件 .env.
	envFile := fmt.Sprintf(".env.%s", env)
	err = godotenv.Overload(envFile)
	if err != nil {
		log.Printf("%s 文件未找到", envFile)
	}
}

func main() {
	setupEnv()

	_, err := dbMysql.ConnectionMysql()
	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认值8080
	}

	r := router.SetupRouter()
	r.Run(":" + port)
}
