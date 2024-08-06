package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/controller"
	"github.com/go-ecommerce-backend-api/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", controller.NewUserController().GetUserID)
	}

	return r
}