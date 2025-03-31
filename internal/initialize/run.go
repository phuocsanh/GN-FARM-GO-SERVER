package initialize

import (
	"fmt"
	"go_ecommerce/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()

	global.Logger.Debug("config log ok", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()

	r:=InitRouter()
	return r
	// r.Run(":8002")
}