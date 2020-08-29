package validations

import (
	"fmt"

	"github.com/syukri21/Paperid-Golang-Testcase/src/validations/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Signup ...
func Signup(c *gin.Context) {
	var errors []map[string]interface{}
	var user schemas.Signup
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	signupValidate := &schemas.Signup{
		Password: user.Password,
		Email:    user.Email,
	}

	Validate(signupValidate, errors)
}

// Signin ...
func Signin(c *gin.Context) {
	var errors []map[string]interface{}
	var user schemas.Signin
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	signinValidate := &schemas.Signin{
		Password: user.Password,
		Email:    user.Email,
	}

	Validate(signinValidate, errors)
}
