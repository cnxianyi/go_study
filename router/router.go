package router

import (
	"go_study/controllers/grammargo"
	testRouter "go_study/controllers/test"
	userRouter "go_study/controllers/user"
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
	_testRouter := r.Group("/test")
	{
		_testRouter.GET("/", testRouter.TestRouter)
	}

	_userRouter := r.Group("/user")
	{
		_userRouter.GET("/", userRouter.UserRouter)
		_userRouter.POST("/add", userRouter.CreateUser)
	}

	grammargoRouter := r.Group("grammargo")
	{
		grammargoRouter.GET("/", grammargo.GrammarGo)

		grammargoRouter.GET("/ct/CompositeTypeByArray", grammargo.CompositeTypeByArray)
		grammargoRouter.GET("/ct/CompositeTypeBySlice", grammargo.CompositeTypeBySlice)
		grammargoRouter.GET("/ct/CompositeTypeByMap", grammargo.CompositeTypeByMap)
		grammargoRouter.GET("/ct/compositeTypeByStruct", grammargo.CompositeTypeByStruct)
		grammargoRouter.GET("/ct/compositeTypeByJson", grammargo.CompositeTypeByJson)
		grammargoRouter.GET("/ct/compositeTypeByTextTemplate", grammargo.CompositeTypeByTextTemplate)

	}

	return r
}
