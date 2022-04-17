package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/mFsl16/clockify-clone/config"
)

func main() {

	app := config.NewApp()
	app.E.Use(middleware.Logger())

	app.Handle()
	app.E.Logger.Fatal(app.E.Start(":8080"))
}
