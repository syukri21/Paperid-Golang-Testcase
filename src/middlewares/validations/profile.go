package validations

import (
	"fmt"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ProfileUpdate ...
func ProfileUpdate(c *gin.Context) {
	var errors []map[string]interface{}
	var data schemas.ProfileUpdate
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	Validate(data, errors)
}