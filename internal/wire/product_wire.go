package wire

import (
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/service/impl"
)

// InitProductService khởi tạo service product
func InitProductService() {
	service.InitProductManagement(impl.NewProductService())
} 