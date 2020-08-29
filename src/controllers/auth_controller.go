package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
	"github.com/syukri21/Paperid-Golang-Testcase/src/validations/schemas"
)

// AuthController -> AuthController
type AuthController struct {
	AuthService service.AuthService
}

// AuthControllerInstance -> AuthControllerInstance
func AuthControllerInstance() AuthController {
	return AuthController{AuthService: service.AuthServiceInstance()}
}

// Signup -> Signup
func (c *AuthController) Signup(ctx *gin.Context) {
	var user entity.User
	ctx.ShouldBindBodyWith(&user, binding.JSON)

	newUser := c.AuthService.Signup(user)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flags.SignupSuccess.Message,
		"data":    newUser,
		"error":   nil,
	})
}

// Signin ...
func (c *AuthController) Signin(ctx *gin.Context) {
	var params schemas.Signin
	ctx.ShouldBindBodyWith(&params, binding.JSON)

	user := c.AuthService.Signin(params)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flags.SigninSuccess.Message,
		"data":    user,
		"error":   nil,
	})
}

// Signout ...
func (c *AuthController) Signout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": flags.SigninSuccess.Message,
		"data":    nil,
		"error":   nil,
	})
}
