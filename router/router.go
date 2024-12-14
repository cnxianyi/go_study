package router

import (
	testRouter "go_study/controllers/test"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置模式
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	// 设置信任代理
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 默认路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	// 路由分组
	test := r.Group("/test")
	{
		test.GET("/", testRouter.TestRouter)
	}

	return r
}
