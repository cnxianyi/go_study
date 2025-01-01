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
创建表
*/
func CreateTable(c *gin.Context) {
	type TableForm struct {
		Table string `form:"table" binding:"required,min=2"` // 表名* 长度不小于2
	}

	var form TableForm

	// 绑定成功 返回nil
	// 绑定失败 返回error
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
		id INT AUTO_INCREMENT PRIMARY KEY,
		usr VARCHAR(255) NOT NULL
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

/*
插入表
*/
func InsertToTable(c *gin.Context) {

	db := mysql.GetDB()

	type tableForm struct {
		Table string `form:"table" binding:"required,min=2"`
		Usr   string `form:"usr" binding:"required,min=2"`
	}

	var form tableForm
	if c.ShouldBind(&form) != nil {
		c.JSON(400, gin.H{
			"error": "绑定失败",
		})
		return
	}

	query := fmt.Sprintf("INSERT INTO %s (usr) VALUES (?)", form.Table)

	res, err := db.Exec(query, form.Usr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("插入失败: %v", err.Error()),
		})
		return
	}

	// 获取最后插入的id
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("获取id失败: %v", err.Error()),
		})
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("插入成功. ID: %d", id),
	})

}
