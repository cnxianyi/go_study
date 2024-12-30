package grammargin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
HTTP2 server 推送
JSONP
PureJson 特定字符编码
SecureJSON 防止 json 劫持
ProtoBuf
*/

/*
AsciiJSON
转义字符
如 < 转义成 \u003C
*/
func AsciiJSON(c *gin.Context) {
	data := gin.H{
		"lang": "Go",   // Go
		"tag":  "<br>", // \u003Cbr\u003E
	}

	c.AsciiJSON(http.StatusOK, data)
}

/*
 1. router中加载file.html模版
 2. 传递模版
 3. html设置

<body>
{{ .title }}
</body>
*/
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "file.html", gin.H{
		"title": "GoLang",
	})
}

/*
Writer
*/
func Writer(c *gin.Context) {
	c.Writer.Write([]byte(`<div>this is html</div>`))
}

/*
Multipart/Urlencoded 绑定
*/
func Form1(c *gin.Context) {

	// binding:"required" 必须值
	type loginForm struct {
		User     string `form:"user" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var form loginForm
	// ShouldBind 自动处理 JSON / FORM
	// ShouldBindJSON 固定 JSON
	// ShouldBindWith(&form, binding.Form) 固定FORM

	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{
				"message": "success",
				"token":   "success",
			})
		}
	} else {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
	}
}

/*
接收表单数据
*/
func Form2(c *gin.Context) {
	// 接收 form 的 message
	message := c.PostForm("message")
	// form 默认值
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"message": message,
		"nick":    nick,
	})
}

/*
Query
?q=1
*/
func Query1(c *gin.Context) {
	q := c.Query("q")

	c.JSON(200, gin.H{
		"q": q,
	})
}

/*
XML
*/
func Xml(c *gin.Context) {
	c.XML(200, gin.H{"message": "hey", "status": http.StatusOK})
	//	<map>
	//	    <message>hey</message>
	//	    <status>200</status>
	// </map>
}

/*
结构体JSON
*/
func Json(c *gin.Context) {
	// 格式化 Name 为 user.
	// 此时 JSON中显示为 user
	type user struct {
		Name string `json:"user"`
		Age  int
	}

	var u1 user
	u1.Name = "Ilya"
	u1.Age = 18

	c.JSON(200, u1)
}

/*
YAML
*/
func Yaml(c *gin.Context) {
	c.YAML(200, gin.H{
		"PORT": 7890,
	})
}
