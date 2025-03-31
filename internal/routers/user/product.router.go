package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{

}
func (r *ProductRouter) InitProductRouter(Router *gin.RouterGroup){
// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("search")
		productRouterPublic.GET("/detail/:id")

	}
}