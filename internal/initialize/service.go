package initialize

import (
	"go_ecommerce/global"
	"go_ecommerce/internal/database"
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/service/impl"
)

func InitService() {
	queries := database.New(global.Mdbc)
	// User serive interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	
	// Add product service initialization
	service.InitProductManagement(impl.NewProductService())
}
