package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "ready",
			"database": "error",
		})
	})
}
