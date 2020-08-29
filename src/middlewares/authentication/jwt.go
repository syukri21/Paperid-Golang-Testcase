package authentication

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/syukri21/Paperid-Golang-Testcase/src/repositories"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
)

type header struct {
	Authorization string
}

// JWT ...
func JWT(ctx *gin.Context) {
	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		jwtError(err)
	}

	if len(h.Authorization) <= 0 {
		jwtError(errors.New("No Authorization Header"))
	}

	tokenString := strings.Replace(h.Authorization, "Bearer ", "", 1)
	jwtClaims, err := verifyToken(tokenString)

	if err != nil {
		jwtError(errors.New("Token Invalid"))
	}

	id := jwtClaims.(jwt.MapClaims)["id"].(string)

	URepo := repositories.UserRepositoryInstance()
	user, err := URepo.GetUserByID(id)

	if err != nil {
		jwtError(errors.New("Token Invalid"))
	}

	if *user.JwtToken != tokenString {
		jwtError(errors.New("Token Invalid"))
	}

	ctx.Set("user", user)
	ctx.Next()
}

func jwtError(err error) {
	errors := []map[string]interface{}{}
	errors = append(errors, map[string]interface{}{
		"message": fmt.Sprint(err),
		"flag":    "UNAUTHORIZED",
	})
	exception.UnAuthorizedRequest("UnAuthorized", errors)
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
