package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v5/middleware"
	"l2.18/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		Skipper:   skipper,
		BeforeNextFunc: func(c *echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))
	controllers.EventController(e)

	port := flag.Int("port", 8080, "app port")

	if err := e.Start(":" + strconv.Itoa(*port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
