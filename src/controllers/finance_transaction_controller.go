package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
)

// FinanceTransactionController ...
type FinanceTransactionController struct {
	FinanceTransactionService service.FinanceTransactionService
}

// FinanceTransactionControllerInstance ...
func FinanceTransactionControllerInstance() FinanceTransactionController {
	return FinanceTransactionController{
		FinanceTransactionService: service.FinanceTransactionServiceInstance(),
	}
}

// Create ...
func (c *FinanceTransactionController) Create(ctx *gin.Context) {

	body := schemas.FinanceTransactionCreate{}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	user := ctx.MustGet("user").(entity.User)

	results := c.FinanceTransactionService.Create(entity.FinanceTransaction{
		Name:             body.Name,
		Description:      body.Description,
		Amount:           body.Amount,
		UserID:           user.ID,
		FinanceAccountID: body.FinanceAccountID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// GetAll ...
func (c *FinanceTransactionController) GetAll(ctx *gin.Context) {

	pagination := schemas.Pagination{}
	ctx.ShouldBindQuery(&pagination)

	results := c.FinanceTransactionService.GetAll(pagination)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// GetByID ...
func (c *FinanceTransactionController) GetByID(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	results := c.FinanceTransactionService.GetByID(ID.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// Delete ...
func (c *FinanceTransactionController) Delete(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	results := c.FinanceTransactionService.Delete(ID.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// Update ...
func (c *FinanceTransactionController) Update(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	body := schemas.FinanceTransactionUpdate{}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	results := c.FinanceTransactionService.Update(ID.ID, entity.FinanceTransaction{
		Name:        body.Name,
		Description: body.Description,
		Amount:      body.Amount,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}
