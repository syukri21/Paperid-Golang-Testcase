package validations

import (
	"fmt"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FinanceAccountTypeCreate ...
func FinanceAccountTypeCreate(c *gin.Context) {
	var errors []map[string]interface{}
	var data schemas.FinanceAccountTypeCreate
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	validate := &schemas.FinanceAccountTypeCreate{
		Name:        data.Name,
		Description: data.Description,
	}

	Validate(validate, errors)
}
