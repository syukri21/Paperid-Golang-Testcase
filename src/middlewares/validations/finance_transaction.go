package validations

import (
	"fmt"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FinanceTransactionCreate ...
func FinanceTransactionCreate(c *gin.Context) {
	var errors []map[string]interface{}
	var data schemas.FinanceTransactionCreate
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	Validate(data, errors)
}

// FinanceTransactionUpdate ...
func FinanceTransactionUpdate(c *gin.Context) {
	var errors []map[string]interface{}
	var data schemas.FinanceTransactionUpdate
	if err := c.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	Validate(data, errors)
}
