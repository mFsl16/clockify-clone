package controller

import "github.com/labstack/echo/v4"

type TaskControllerImpl struct {
}

func NewTaskController() TaskController {
	return *TaskControllerImpl{}
}

func (controller *TaskControllerImpl) AddTask(c echo.Context) error {

	return nil
}
