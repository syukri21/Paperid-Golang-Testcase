package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
)

// FinanceAccountController ...
type FinanceAccountController struct {
	FinnanceAccountService service.FinanceAccountService
}

// FinanceAccountControllerInstance ...
func FinanceAccountControllerInstance() FinanceAccountController {
	return FinanceAccountController{
		FinnanceAccountService: service.FinanceAccountServiceInstance(),
	}
}

// GetAll ...
func (c *FinanceAccountController) GetAll(ctx *gin.Context) {

	pagination := schemas.Pagination{}
	ctx.ShouldBindQuery(&pagination)

	results := c.FinnanceAccountService.GetAll(pagination)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})

}
