package controllers

import (
	"net/http"

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

	results := c.FinnanceAccountService.GetAll(pagination)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// GetByID ...
func (c *FinanceAccountController) GetByID(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	results := c.FinnanceAccountService.GetByID(ID.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// Update ...
func (c *FinanceAccountController) Update(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	body := schemas.FinanceAccountUpdate{}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	results := c.FinnanceAccountService.Update(ID.ID, entity.FinanceAccount{
		Name:        body.Name,
		Description: body.Description,
		TypeID:      body.TypeID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    results,
		"error":   nil,
	})
}

// Delete ...
func (c *FinanceAccountController) Delete(ctx *gin.Context) {

	ID := schemas.ID{}
	ctx.ShouldBindUri(&ID)

	results := c.FinnanceAccountService.Delete(ID.ID)

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
