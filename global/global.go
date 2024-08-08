package global

import (
	"github.com/go-ecommerce-backend-api/pkg/logger"
	"github.com/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	MDb *gorm.DB
)