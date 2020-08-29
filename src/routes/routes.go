package routes

import (
	"net/http"

	"github.com/syukri21/Paperid-Golang-Testcase/src/controllers"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/authentication"
	"github.com/syukri21/Paperid-Golang-Testcase/src/validations"

	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {
	// Auth
	{
		controller := controllers.AuthControllerInstance()
		g.POST("/signup", validations.Signup, controller.Signup)
		g.POST("/signin", validations.Signin, controller.Signin)
		g.GET("/signout", authentication.JWT, controller.Signout)
	}

	// Check Health
	{
		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "ready",
				"database": "error",
			})
		})
	}
}
