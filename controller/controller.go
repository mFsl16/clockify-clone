package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"github.com/mFsl16/clockify-clone/usecase"
	"github.com/sirupsen/logrus"
)

type TaskControllerImpl struct {
	E *echo.Echo
	U usecase.Usecase
}

func (controller *TaskControllerImpl) Handle() {
	controller.E.GET("/", controller.Hello)
	controller.E.POST("/v1/task", controller.AddTask)
	controller.E.POST("/v1/project", controller.addProject)
}

func NewTaskController(e *echo.Echo, usecase usecase.Usecase) *TaskControllerImpl {
	return &TaskControllerImpl{
		E: e,
		U: usecase,
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

	logrus.Info("[Receive Request:", requestBody, "]")

	task := controller.U.AddTask(c.Request().Context(), requestBody)

	return c.JSON(http.StatusOK, task)
}

func (controller *TaskControllerImpl) addProject(c echo.Context) error {

	requestBody := request.ProjectRq{}
	err := c.Bind(&requestBody)
	if err != nil {
		panic(err)
	}

	logrus.Info("[Receive Request:", requestBody, "]")

	project := controller.U.AddProject(c.Request().Context(), requestBody)
	return c.JSON(http.StatusOK, project)
}
