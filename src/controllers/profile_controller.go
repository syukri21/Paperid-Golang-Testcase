package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/syukri21/Paperid-Golang-Testcase/src/database/entity"
	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/validations/schemas"
	"github.com/syukri21/Paperid-Golang-Testcase/src/service"
)

// ProfileController ...
type ProfileController struct {
	ProfileService service.ProfileService
	UserService    service.UserService
}

// ProfileControllerInstance ...
func ProfileControllerInstance() ProfileController {
	return ProfileController{
		ProfileService: service.ProfileServiceInstance(),
		UserService:    service.UserServiceInstance(),
	}
}

// Update ...
func (c *ProfileController) Update(ctx *gin.Context) {

	body := schemas.ProfileUpdate{}
	ctx.ShouldBindBodyWith(&body, binding.JSON)

	user := ctx.MustGet("user").(entity.User)

	c.ProfileService.UpdateOrCreate(entity.Profile{
		Address:     body.Address,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		PhoneNumber: body.PhoneNumber,
		UserID:      user.ID,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    nil,
		"error":   nil,
	})
}

// GetByID ...
func (c *ProfileController) GetByID(ctx *gin.Context) {

	user := ctx.MustGet("user").(entity.User)

	profile := c.UserService.GetOne(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    profile,
		"error":   nil,
	})
}
