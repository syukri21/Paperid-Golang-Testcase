package controllers

import (
	"net/http"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
)

// FinanceAccountTypeController -> FinanceAccountTypeController
type FinanceAccountTypeController struct {
	FinanceAccountTypeService service.FinanceAccountTypeService
}

// FinanceAccountTypeControllerInstance -> FinanceAccountTypeControllerInstance
func FinanceAccountTypeControllerInstance() FinanceAccountTypeController {
	return FinanceAccountTypeController{
		FinanceAccountTypeService: service.FinanceAccountTypeServiceInstance(),
	}
}

// Create ...
// POST
func (c *FinanceAccountTypeController) Create(ctx *gin.Context) {
	params := schemas.FinanceAccountTypeCreate{}
	ctx.ShouldBindBodyWith(&params, binding.JSON)
	data := c.FinanceAccountTypeService.Create(params)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}

// GetAll ...
// GET
func (c *FinanceAccountTypeController) GetAll(ctx *gin.Context) {
	data := c.FinanceAccountTypeService.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}

// GetOne ...
// GET/:id
func (c *FinanceAccountTypeController) GetOne(ctx *gin.Context) {
	var TypeID schemas.TypeID
	ctx.ShouldBindUri(&TypeID)
	data := c.FinanceAccountTypeService.GetByID(TypeID.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}

// Update ...
// PUT/:id
func (c *FinanceAccountTypeController) Update(ctx *gin.Context) {
	var TypeID schemas.TypeID
	ctx.ShouldBindUri(&TypeID)

	var body schemas.FinanceAccountTypeUpdate
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	data := c.FinanceAccountTypeService.Update(TypeID.ID, entity.FinanceAccountType{
		Name:        body.Name,
		Description: body.Description,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}

// Delete ...
// DELETE/:id
func (c *FinanceAccountTypeController) Delete(ctx *gin.Context) {
	var TypeID schemas.TypeID
	ctx.ShouldBindUri(&TypeID)
	data := c.FinanceAccountTypeService.Delete(TypeID.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
		"error":   nil,
	})
}
