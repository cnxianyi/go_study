package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 配置env
func SetupEnv() {
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
