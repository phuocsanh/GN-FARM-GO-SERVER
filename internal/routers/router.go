package routers

// import (
// 	// c "go_ecommerce/internal/controlller"
// 	"go_ecommerce/internal/routers"
// 	"go_ecommerce/middlewares"

// 	"github.com/gin-gonic/gin"
// )

// func NewRouter() *gin.Engine {
// r := gin.Default()
// r.Use(middlewares.AuthenMiddleware())

// managerRouter := routers.RouterGroupApp.Manager
// 	userRouter := routers.RouterGroupApp.User

// 	MainGroup := r.Group("/v1")
// 	{
// 		MainGroup.GET("/check_status") // tracking monitor
// 	}
// 	{
// 		managerRouter.InitUserRouter(MainGroup)
// 		managerRouter.InitAdminRouter(MainGroup)
// 	}
// 	{
// 		userRouter.InitUserRouter(MainGroup)
// 		userRouter.InitProductRouter(MainGroup)
// 	}
// return r
// }

