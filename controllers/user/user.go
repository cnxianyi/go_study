package user

import (
	"fmt"
	"go_study/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Context) {
	c.String(http.StatusOK, "user")
}

type CreateUserReq struct {
	User     string `json:"user" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {

	var req CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := user.AddUser(req.User, req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": fmt.Sprintf("用户 %d 添加成功", id),
		})
	}
}
