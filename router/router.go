package router

import (
	"go_study/controllers/grammargo"
	"go_study/controllers/test"
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
	testRouter := r.Group("/test")
	{
		testRouter.GET("/", test.TestRouter)
	}

	grammargoRouter := r.Group("grammargo")
	{
		grammargoRouter.GET("/", grammargo.GrammarGo)
	}

	return r
}
