package grammarmysql

import (
	"fmt"
	"go_study/database/mysql"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	db := mysql.GetDB()

	type formTable struct {
		Table string `form:"table" binding:"required,min=2"`
	}

	var form formTable

	if c.ShouldBind(&form) != nil {
		c.JSON(200, gin.H{
			"error": "格式错误",
		})
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s`, form.Table)
	// db.Query 会返回 *sql.Rows 对象 : 结果集
	// 不同于 Exec. *sql.Rows 需要正确关闭,以释放数据库连接资源
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(200, gin.H{
			"error": "查询错误",
		})
		return
	}
	// 关闭连接
	defer rows.Close()

	// 创建结果集 结构体
	type User struct {
		ID  int    `json:"id"`
		Usr string `json:"usr"`
	}

	// 创建结果集 切片
	var res []User

	// 调用结果集.Next() 方法. 会将指针向下移动. 此时rows指向下一行的结果.
	// 没有更多行结果时. 循环结束
	for rows.Next() {
		var user User
		// 扫描结果. 将id 和 usr 扫描到 user 变量中
		if err := rows.Scan(&user.ID, &user.Usr); err != nil {
			c.JSON(400, gin.H{
				"error": "扫描错误",
			})
			return
		}

		// 将扫描结果 附加到res
		res = append(res, user)
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}
