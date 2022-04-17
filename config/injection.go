//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/controller"
	"github.com/mFsl16/clockify-clone/repository"
	"github.com/mFsl16/clockify-clone/usecase"
)

func NewApp() *controller.TaskControllerImpl {
	wire.Build(
		echo.New,
		controller.NewTaskController,
		usecase.NewUsecase,
		repository.NewDatabase,
		repository.NewProjectRepository,
		repository.NewTaskRepository,
	)

	return nil
}
