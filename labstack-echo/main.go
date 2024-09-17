package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	setUpMiddleware(e)
	setUpRouting(e)
	e.Logger.Fatal(e.Start(":4200"))
}
