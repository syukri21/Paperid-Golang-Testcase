package controllers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	logrus.Info(pagination)

	results := c.FinnanceAccountService.GetAll(pagination)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// Create ...
func (c *FinanceAccountController) Create(ctx *gin.Context) {

	body := schemas.FinanceAccountCreate{}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	user := ctx.MustGet("user").(entity.User)

	results := c.FinnanceAccountService.Create(entity.FinanceAccount{
		Name:        body.Name,
		Description: body.Description,
		TypeID:      body.TypeID,
		UserID:      user.ID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}
