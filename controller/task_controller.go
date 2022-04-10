package controller

import "github.com/labstack/echo/v4"

type TaskController interface {
	AddTask(c echo.Context) error
}
