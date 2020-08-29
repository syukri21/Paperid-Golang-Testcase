package authorization

import (
	"errors"
	"fmt"

	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"

	"github.com/gin-gonic/gin"
)

// Role ...
func Role(roles ...string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		usermap, exist := ctx.Get("user")
		if !exist {
			notFoudErr(errors.New("Not Found"))
		}

		user := usermap.(entity.User)
		match := include(user.Role, roles)
		if !match {
			notFoudErr(errors.New("UnAuthorized"))
		}
		ctx.Next()
	}
}

func include(role string, roles []string) (match bool) {

	for _, r := range roles {
		if r == role {
			match = true
			break
		} else {
			match = false
		}
	}

	return
}

func notFoudErr(err error) {
	errors := []map[string]interface{}{}
	errors = append(errors, map[string]interface{}{
		"message": fmt.Sprint(err),
		"flag":    "UNAUTHORIZED",
	})
	exception.NotFound("UnAuthorized", errors)
}
