package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mFsl16/clockify-clone/config"
)

func main() {
	e := config.NewApp()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "This is clockify api clone")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
