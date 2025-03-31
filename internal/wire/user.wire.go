//go:build wireinject

package wire

import (
	"go_ecommerce/internal/controlller"
	"go_ecommerce/internal/repo"
	"go_ecommerce/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler()(*controlller.UserController, error)  {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controlller.NewUserController,
	)
	return new(controlller.UserController), nil
}