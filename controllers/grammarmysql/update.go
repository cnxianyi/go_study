package grammarmysql

import (
	"fmt"
	"go_study/database/mysql"

	"github.com/gin-gonic/gin"
)

func UpdateToTable(c *gin.Context) {

	db := mysql.GetDB()

	type formData struct {
		Table string `form:"table" binding:"required,min=2"`
		Id    int    `form:"id" binding:"required"`
		Usr   string `form:"usr" binding:"required,min=2"`
	}

	var form formData

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{
			// Error() 将err转为字符串
			"error": err.Error(),
		})
		return
	}

	query := fmt.Sprintf(`UPDATE %s SET usr = %s WHERE id = %d`, form.Table, form.Usr, form.Id)

	_, err := db.Exec(query)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update success",
	})

}
