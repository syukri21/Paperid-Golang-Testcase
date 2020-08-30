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

	// Finance Account
	{
		controller := controllers.FinanceAccountControllerInstance()
		g.POST(
			"/finance-accounts",
			authentication.JWT,
			validations.FinanceAccountCreate,
			controller.Create,
		)
		g.GET("/finance-accounts", authentication.JWT, controller.GetAll)
		g.GET("/finance-accounts/:id", authentication.JWT, controller.GetByID)
		g.DELETE("/finance-accounts/:id", authentication.JWT, controller.Delete)
		g.PUT(
			"/finance-accounts/:id",
			authentication.JWT,
			validations.FinanceAccountUpdate,
			controller.Update,
		)
	}

	// Finance Transaction
	{
		controller := controllers.FinanceTransactionControllerInstance()
		g.POST(
			"/finance-transactions",
			authentication.JWT,
			validations.FinanceTransactionCreate,
			controller.Create,
		)
		g.GET("/finance-transactions", authentication.JWT, controller.GetAll)
		g.GET("/finance-transactions/:id", authentication.JWT, controller.GetByID)
		g.DELETE("/finance-transactions/:id", authentication.JWT, controller.Delete)
		g.PUT(
			"/finance-transactions/:id",
			authentication.JWT,
			validations.FinanceTransactionUpdate,
			controller.Update,
		)
	}

	// Finance Account Type
	{
		controller := controllers.FinanceAccountTypeControllerInstance()
		g.POST(
			"/finance-account/types",
			authentication.JWT,
			authorization.Role(authorization.ADMIN),
			validations.FinanceAccountTypeCreate,
			controller.Create,
		)
		g.GET("/finance-account/types", authentication.JWT, controller.GetAll)
		g.GET("/finance-account/types/:id", authentication.JWT, controller.GetOne)
		g.DELETE("/finance-account/types/:id", authentication.JWT, authorization.Role(authorization.ADMIN), controller.Delete)
		g.PUT(
			"/finance-account/types/:id",
			authentication.JWT,
			authorization.Role(authorization.ADMIN),
			validations.FinanceAccountTypeUpdate,
			controller.Update,
		)
	}

	// Check Health
	{
		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "ready",
			})
		})
	}
}
