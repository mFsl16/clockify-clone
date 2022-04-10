package main

import "github.com/mFsl16/clockify-clone/config"

func main() {

	app := config.NewApp()

	app.Handle()
	app.E.Logger.Fatal(app.E.Start(":8080"))
}
