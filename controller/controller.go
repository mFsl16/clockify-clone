package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/model"
)

type TaskControllerImpl struct {
	E *echo.Echo
}

func (controller *TaskControllerImpl) Handle() {
	controller.E.GET("/", controller.Hello)
	controller.E.POST("/v1/task", controller.AddTask)
	controller.E.POST("/v1/project", controller.addProject)
}

func NewTaskController(e *echo.Echo) *TaskControllerImpl {
	return &TaskControllerImpl{
		E: e,
	}
}

func (controller *TaskControllerImpl) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "This is clockify api clone")
}

func (controller *TaskControllerImpl) AddTask(c echo.Context) error {

	requestBody := model.Task{}
	err := c.Bind(&requestBody)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, requestBody)
}

func (controller *TaskControllerImpl) addProject(c echo.Context) error {

	requestBody := model.Project{}
	err := c.Bind(&requestBody)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, requestBody)
}
