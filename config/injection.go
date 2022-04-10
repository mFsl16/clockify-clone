//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/controller"
)

func NewApp() *controller.TaskControllerImpl {
	wire.Build(
		echo.New,
		controller.NewTaskController,
	)

	return nil
}
