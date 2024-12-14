package grammargo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GrammarGo(c *gin.Context) {
	c.String(http.StatusOK, "grammar for go")
}
