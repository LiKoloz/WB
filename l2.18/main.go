package main

import (
	"flag"
	"strconv"

	"github.com/labstack/echo"
	"l2.18/controllers"
)

func main() {
	e := echo.New()

	controllers.EventController(e)

	port := flag.Int("port", 8080, "app port")

	if err := e.Start(":" + strconv.Itoa(*port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
