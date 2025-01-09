package router

import (
	"go_study/controllers/grammargin"
	"go_study/controllers/grammargo"
	"go_study/controllers/grammarmysql"
	"go_study/controllers/grammarredis"
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

	// 加载HTML模板
	// r.LoadHTMLFiles("templates/gin/file.html")
	// or
	r.LoadHTMLGlob("templates/gin/file.html")

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

	// Go语法
	grammargoRouter := r.Group("grammargo")
	{
		grammargoRouter.GET("/", grammargo.GrammarGo)

		grammargoRouter.GET("/bt/basicTypes", grammargo.BasicTypes)

		grammargoRouter.GET("/ct/CompositeTypeByArray", grammargo.CompositeTypeByArray)
		grammargoRouter.GET("/ct/CompositeTypeBySlice", grammargo.CompositeTypeBySlice)
		grammargoRouter.GET("/ct/CompositeTypeByMap", grammargo.CompositeTypeByMap)
		grammargoRouter.GET("/ct/compositeTypeByStruct", grammargo.CompositeTypeByStruct)
		grammargoRouter.GET("/ct/compositeTypeByJson", grammargo.CompositeTypeByJson)
		grammargoRouter.GET("/ct/compositeTypeByTextTemplate", grammargo.CompositeTypeByTextTemplate)

		grammargoRouter.GET("/fc/FuncGrammar", grammargo.FuncGrammar)

		// 方法
		grammargoRouter.GET("/me/MethodTest", grammargo.MethodTest)
		grammargoRouter.GET("/me/ExpandStruct", grammargo.ExpandStruct)

		// 接口
		grammargoRouter.GET("/in/interfaceTest", grammargo.InterfaceTest)

		// 协程
		grammargoRouter.GET("/gr/Goroutines", grammargo.GoroutinesTest)

		// 反射
		grammargoRouter.GET("/rf/reflect", grammargo.ReflectTest)

		// gin 渲染
		grammargoRouter.GET("/gin/AsciiJSON", grammargin.AsciiJSON)
		grammargoRouter.GET("/gin/html", grammargin.Html)      // html模版
		grammargoRouter.POST("/gin/form1", grammargin.Form1)   // 表单
		grammargoRouter.POST("/gin/form2", grammargin.Form2)   // 接收表单数据
		grammargoRouter.POST("/gin/query1", grammargin.Query1) // 接收 Query

		grammargoRouter.GET("/gin/xml", grammargin.Xml)   // XML格式
		grammargoRouter.GET("/gin/json", grammargin.Json) // XML格式
		grammargoRouter.GET("/gin/yaml", grammargin.Yaml) // yaml格式
	}

	// Mysql语法
	grammarmysqlRouter := r.Group("grammarmysql")
	{
		grammarmysqlRouter.GET("/add/add", grammarmysql.Add)

		// 创建表
		grammarmysqlRouter.POST("/create/createTable", grammarmysql.CreateTable)
		// 插入数据
		grammarmysqlRouter.POST("/create/InsertToTable", grammarmysql.InsertToTable)

		// 更新数据
		grammarmysqlRouter.POST("/update/updateToTable", grammarmysql.UpdateToTable)

		// 获取数据
		grammarmysqlRouter.POST("/read/ReadAll", grammarmysql.ReadAll)

		// 删除数据
		grammarmysqlRouter.POST("/delete/DeleteToTable", grammarmysql.DeleteToTable)
	}

	grammarredisRouter := r.Group("grammarredis")
	{
		grammarredisRouter.POST("/redis/redistest", grammarredis.RedisTest)

		// string
		grammarredisRouter.GET("/redis/stringtest", grammarredis.StringTest)

		// list
		grammarredisRouter.GET("/redis/listtest", grammarredis.ListTest)

		// set
		grammarredisRouter.GET("/redis/settest", grammarredis.SetTest)
	}

	return r
}
