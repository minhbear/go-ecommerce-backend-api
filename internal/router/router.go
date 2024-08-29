package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	return r
}
