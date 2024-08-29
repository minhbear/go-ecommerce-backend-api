package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/service"
	"github.com/go-ecommerce-backend-api/internal/vo"
	"github.com/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrarRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrInvalidParams, err.Error())
		return
	}
	result := userController.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
