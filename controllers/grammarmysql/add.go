package grammarmysql

import (
	"fmt"
	"go_study/database/mysql"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

/*
新增表
*/
func AddTable(c *gin.Context) {
	type TableForm struct {
		Table string `form:"table" binding:"required,min=2"` // 表名* 长度不小于2
	}

	var form TableForm

	if c.ShouldBind(&form) != nil { // 检查绑定是否成功
		c.JSON(400, gin.H{
			"error": "绑定失败",
		})
		return
	}

	db := mysql.GetDB()

	// 创建表时 不支持 ? 占位符
	// 因此选择使用 字符串拼接
	//! 生产环境应注意 SQL注入
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY
	)
	`, form.Table)
	base, err := db.Exec(query)
	if err != nil {
		fmt.Printf("创建表失败: %v", err)
		c.JSON(400, gin.H{
			"error": fmt.Errorf("创建表失败: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"table":   form.Table,
		"a":       base,
	})

}
