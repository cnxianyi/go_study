package grammargo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GrammarGo(c *gin.Context) {
	stateGrammar()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
