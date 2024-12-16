package grammargo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GrammarGo(c *gin.Context) {
	stateGrammar()
	basicTypes()

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
}
