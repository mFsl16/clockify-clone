package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mFsl16/clockify-clone/exception"
	"github.com/mFsl16/clockify-clone/model/request"
	"github.com/mFsl16/clockify-clone/usecase"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	controller.E.PUT("/v1/project/:id", controller.UpdateProject)
	controller.E.DELETE("/v1/task/:id", controller.DeleteTask)
	controller.E.DELETE("/v1/project/:id", controller.DeleteProject)
}

func NewController(e *echo.Echo, usecase usecase.Usecase) *Controller {
	controller := Controller{
		E: e,
		U: usecase,
	}

	controller.E.Use(middleware.Recover())
	controller.E.HTTPErrorHandler = exception.CustomHttpErrorHandler
	return &controller
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

	logrus.WithFields(
		logrus.Fields{
			"requestbody": taskUpdate,
			"id":          id,
		}).Info("[RECEIVE REQUEST]")

	if errBindRq != nil || errBindParam != nil {
		panic("error bind request: " + errBindRq.Error() + " " + errBindParam.Error())
	}

	taskResult := controller.U.UpdateTask(c.Request().Context(), id, taskUpdate)

	return c.JSON(http.StatusOK, taskResult)
}

func (controller *Controller) UpdateProject(c echo.Context) error {

	projectUpdate := request.ProjectRq{}
	errBindRq := c.Bind(&projectUpdate)

	id, errBindParam := strconv.Atoi(c.Param("id"))

	if errBindParam != nil || errBindRq != nil {
		panic("error bind request: " + errBindRq.Error() + " " + errBindParam.Error())
	}

	logrus.WithFields(logrus.Fields{
		"requestbody": projectUpdate,
		"id":          id,
	}).Info("[RECEIVE REQUEST]")

	projectResult := controller.U.UpdateProject(c.Request().Context(), id, projectUpdate)

	return c.JSON(http.StatusOK, projectResult)
}

func (controller *Controller) DeleteTask(c echo.Context) error {

	id, errBindParam := strconv.Atoi(c.Param("id"))

	if errBindParam != nil {
		panic("error binding param: " + errBindParam.Error())
	}

	deleteTask := controller.U.DeleteTask(c.Request().Context(), id)

	return c.JSON(http.StatusOK, deleteTask)
}

func (controller *Controller) DeleteProject(c echo.Context) error {

	id, errBindParam := strconv.Atoi(c.Param("id"))

	if errBindParam != nil {
		panic("error bind param: " + errBindParam.Error())
	}

	deleteProject := controller.U.DeleteProject(c.Request().Context(), id)

	return c.JSON(http.StatusOK, deleteProject)
}
