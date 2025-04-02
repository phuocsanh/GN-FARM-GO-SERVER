package user

import (
	"go_ecommerce/internal/controlller/product"
	"go_ecommerce/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (r *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// Public routes for product browsing
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("", product.Product.GetAllProducts)
		productRouterPublic.GET("/:id", product.Product.GetProductByID)
		productRouterPublic.GET("/search", product.Product.SearchProducts)
		productRouterPublic.GET("/discounts", product.Product.GetProductsByDiscount)
		productRouterPublic.GET("/bestsellers", product.Product.GetProductsBySelled)
	}

	// Private routes for product management
	productRouterPrivate := Router.Group("/product")
	productRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		// Product CRUD operations
		productRouterPrivate.POST("/create", product.Product.CreateProduct)
		productRouterPrivate.PUT("/update/:id", product.Product.UpdateProduct)
		
		// Product status management
		productRouterPrivate.PUT("/publish/:id", product.Product.PublishProduct)
		productRouterPrivate.PUT("/unpublish/:id", product.Product.UnPublishProduct)
		
		// Shop-specific product lists
		productRouterPrivate.GET("/drafts", product.Product.GetAllDraftsForShop)
		productRouterPrivate.GET("/published", product.Product.GetAllPublishForShop)
	}
}
