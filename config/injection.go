//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func NewApp() *echo.Echo {
	wire.Build(
		echo.New,
	)

	return nil
}
