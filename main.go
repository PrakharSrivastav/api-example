// Package main shows api example
package main

import (
	"encoding/json"
	"github.com/PrakharSrivastav/api-example/internal/app"
	"github.com/labstack/echo"
	"io/ioutil"
)

func main() {
	// create Echo
	e := echo.New()
	app := app.New()
	// do server and router configuration
	app.Configure(e)
	app.Middleware(e)
	// add routes
	app.Routes(e)
	// start server

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		err.Error()
	}
	ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.StartServer(app.Server()))
}
