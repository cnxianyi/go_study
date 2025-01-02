package grammarmysql

import (
	"fmt"
	"go_study/database/mysql"

	"github.com/gin-gonic/gin"
)

func DeleteToTable(c *gin.Context) {

	var db = mysql.GetDB()

	type formTable struct {
		Table string `form:"table" binding:"required,min=2"`
		Id    int    `form:"id" binding:"required"`
	}

	var form formTable

	if c.ShouldBind(&form) != nil {
		c.JSON(400, gin.H{
			"error": "数据错误",
		})
		return
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, form.Table, form.Id)

	res, err := db.Exec(query)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r1, _ := res.RowsAffected()

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("删除了 %d 行", r1),
	})
}
