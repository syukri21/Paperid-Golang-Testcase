package routes

import (
	"net/http"

	"github.com/syukri21/Paperid-Golang-Testcase/src/controllers"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/authentication"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/authorization"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations"

	"github.com/gin-gonic/gin"
)

// Router middleware to handler routes
func Router(g *gin.RouterGroup) {

	// Auth
	{
		controller := controllers.AuthControllerInstance()
		g.POST("/signup", validations.Signup, controller.Signup)
		g.POST("/signup/admin", validations.Signup, controller.SignupAdmin)
		g.POST("/signin", validations.Signin, controller.Signin)
		g.GET("/signout", authentication.JWT, authorization.Role(authorization.USER, authorization.ADMIN), controller.Signout)
	}

	// Finance Type
	{
		controller := controllers.FinanceAccountTypeControllerInstance()
		g.POST(
			"/finance-types",
			authentication.JWT,
			authorization.Role(authorization.ADMIN),
			validations.FinanceAccountTypeCreate,
			controller.Create,
		)
		g.GET("/finance-types", authentication.JWT, controller.GetAll)
		g.GET("/finance-types/:id", authentication.JWT, controller.GetOne)
		g.DELETE("/finance-types/:id", authentication.JWT, authorization.Role(authorization.ADMIN), controller.Delete)
		g.PUT("/finance-types/:id", authentication.JWT, authorization.Role(authorization.ADMIN), controller.Update)
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
