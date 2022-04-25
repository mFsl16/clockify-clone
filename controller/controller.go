package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/model/request"
	"github.com/mFsl16/clockify-clone/usecase"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	E *echo.Echo
	U usecase.Usecase
}

func (controller *Controller) Handle() {
	controller.E.GET("/", controller.Hello)
	controller.E.POST("/v1/task", controller.AddTask)
	controller.E.POST("/v1/project", controller.addProject)
	controller.E.GET("/v1/project/:id", controller.GetProjectById)
	controller.E.GET("/v1/task/:id", controller.GetTaskById)
	controller.E.GET("/v1/project", controller.GetAllProject)
	controller.E.GET("/v1/task", controller.GetAllTasks)
	controller.E.PUT("/v1/task/:id", controller.UpdateTask)
}

func NewController(e *echo.Echo, usecase usecase.Usecase) *Controller {
	return &Controller{
		E: e,
		U: usecase,
	}
}

func (controller *Controller) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "This is clockify api clone")
}

func (controller *Controller) AddTask(c echo.Context) error {

	requestBody := request.TaskRq{}
	err := c.Bind(&requestBody)
	if err != nil {
		panic(err)
	}

	logrus.Info("[Receive Request:", requestBody, "]")

	task := controller.U.AddTask(c.Request().Context(), requestBody)

	return c.JSON(http.StatusOK, task)
}

func (controller *Controller) addProject(c echo.Context) error {

	requestBody := request.ProjectRq{}
	err := c.Bind(&requestBody)
	if err != nil {
		panic(err)
	}

	logrus.Info("[Receive Request:", requestBody, "]")

	project := controller.U.AddProject(c.Request().Context(), requestBody)
	return c.JSON(http.StatusOK, project)
}

func (controller *Controller) GetProjectById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("Error parsing id: " + err.Error())
	}

	project := controller.U.GetProjectById(c.Request().Context(), id)

	return c.JSON(http.StatusOK, project)
}

func (controller *Controller) GetTaskById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("Error parsing id: " + err.Error())
	}

	task := controller.U.GetTaskById(c.Request().Context(), id)

	return c.JSON(http.StatusOK, task)
}

func (controller *Controller) GetAllProject(c echo.Context) error {

	projects := controller.U.GetAllProject(c.Request().Context())

	return c.JSON(http.StatusOK, projects)
}

func (controller *Controller) GetAllTasks(c echo.Context) error {

	tasks := controller.U.GetAllTasks(c.Request().Context())

	return c.JSON(http.StatusOK, tasks)
}

func (controller *Controller) UpdateTask(c echo.Context) error {

	taskUpdate := request.TaskRq{}
	errBindRq := c.Bind(&taskUpdate)

	id, errBindParam := strconv.Atoi(c.Param("id"))

	if errBindRq != nil || errBindParam != nil {
		panic("error bind request: " + errBindRq.Error())
	}

	fmt.Println(taskUpdate)

	taskResult := controller.U.UpdateTask(c.Request().Context(), id, taskUpdate)

	return c.JSON(http.StatusOK, taskResult)
}
