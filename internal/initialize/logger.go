package initialize

import (
	"go_ecommerce/global"
	"go_ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
