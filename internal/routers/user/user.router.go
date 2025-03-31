package user

import (
	"go_ecommerce/internal/controlller/account"
	"go_ecommerce/internal/middlewares"

	// "go_ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	/*
		Nếu không dùng dependency injection
		ur:=repo.NewUserRepository()
		us:= service.NewUserService(ur)
		userHandlerNonDependency:=controlller.NewUserController(us)
	*/

	//Dùng dependency injection by wire
	// userController, _:= wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		/*
			Nếu không dùng dependency injection
			userRouterPublic.POST("register", userHandlerNonDependency.Register)
		*/

		//Dùng dependency injection by wire
		// userRouterPublic.POST("/register", userController.Register)

		userRouterPublic.POST("/refresh-token", account.Login.RefreshToken)
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/verify-account", account.Login.VerifyOTP)
		userRouterPublic.POST("/update-pass-register", account.Login.UpdatePasswordRegister)
	}
	// private router
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		userRouterPrivate.GET("get-info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	}
}
