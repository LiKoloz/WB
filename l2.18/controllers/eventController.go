package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"l2.18/models"
	"l2.18/services"
)

// EventController = все методы контроллера
func EventController(e *echo.Echo) {
	e.GET("/day/:day", func(c echo.Context) error {
		day, err := strconv.Atoi(c.Param("day"))
		if err != nil {
			return c.String(400, "Неккоректный день")
		}
		result := services.GetByDay(day)
		return c.JSON(200, result)
	})
	e.GET("/week", func(c echo.Context) error {
		result := services.GetByWeek()
		return c.JSON(200, result)
	})

	e.GET("/month/:month", func(c echo.Context) error {
		month, err := strconv.Atoi(c.Param("month"))
		if err != nil {
			return c.String(400, "Неккоректный месяц")
		}
		result := services.GetByMonth(month)
		return c.JSON(200, result)
	})

	e.POST("/create", func(c echo.Context) error {
		event := new(models.Event)
		if err := c.Bind(event); err != nil {
			return c.String(400, "Uncorrect event")
		}
		go services.AddEvent(*event)
		return c.String(200, "done")
	})

	e.POST("/updete", func(c echo.Context) error {
		event := new(models.Event)
		if err := c.Bind(event); err != nil {
			return c.String(400, "Uncorrect event")
		}
		go services.UpdateEvent(*event)
		return c.String(200, "done")
	})

	e.POST("delete/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(400, "Неккоректный id")
		}
		err = services.DeleteEvent(id)
		if err == nil {
			return c.JSON(200, "")
		}
		return c.String(503, err.Error())
	})
}
