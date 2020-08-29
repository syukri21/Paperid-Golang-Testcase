package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
	"github.com/syukri21/Paperid-Golang-Testcase/src/utils/flags"
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
