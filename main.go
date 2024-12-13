package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 配置env
func setupEnv() {
	// godotenv
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
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

// 配置路由
func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	return r
}

func main() {
	setupEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认值8080
	}

	r := setupRouter()
	r.Run(":" + port)
}
