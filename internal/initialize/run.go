package initialize

import (
	"fmt"

	"github.com/go-ecommerce-backend-api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.MySQL.Username, global.Config.MySQL.Password)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run("8002")
}