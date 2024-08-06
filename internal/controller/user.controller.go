package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/service"
	"github.com/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (userController *UserController) GetUserID(c *gin.Context) {	
	response.SuccessResponse(c, 20001, userController.userService.GetInfoUser())
}