package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct {}

func (userRouter *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController, _ := wire.InitializeUserController()

	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")	 
	}

	userPrivateRouter := r.Group("/user")
	
	{
		userPrivateRouter.GET("/info")
	}
}