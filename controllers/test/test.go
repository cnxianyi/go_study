package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRouter(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
