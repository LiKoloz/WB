package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"l2.18/models"
	"l2.18/services"
)

func AddEventHandler(c echo.Context) error {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	services.AddEvent(event)
	return c.JSON(http.StatusCreated, event)
}

func GetByDayHandler(c echo.Context) error {
	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid day"})
	}
	events := services.GetByDay(day)
	return c.JSON(http.StatusOK, events)
}

func GetByWeekHandler(c echo.Context) error {
	events := services.GetByWeek()
	return c.JSON(http.StatusOK, events)
}

func GetByMonthHandler(c echo.Context) error {
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid month"})
	}
	events := services.GetByMonth(month)
	return c.JSON(http.StatusOK, events)
}

func DeleteEventHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	err = services.DeleteEvent(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
