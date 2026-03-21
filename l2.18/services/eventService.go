package services

import (
	"errors"
	"time"

	"l2.18/models"
)

var chache = make([]models.Event, 0)

// AddEvent добавления события
func AddEvent(event models.Event) {
	chache = append(chache, event)
}

// UpdateEvent обнволение ивента
func UpdateEvent(event models.Event) {
	for i := 0; i < len(chache); i++ {
		if chache[i].Id == event.Id {
			chache[i] = event
		}
	}
}

// DeleteEvent удаление ивента
func DeleteEvent(id int) error {
	var flag = false
	for i := 0; i < len(chache); i++ {
		if chache[i].Id == id {
			flag = true
			chache = append(chache[:i], chache[i+1:]...)
		}
	}
	if flag == true {
		return nil
	}
	return errors.New("Такого пользователя не существует")
}

// GetByDay получение событий по дню
func GetByDay(day int) []models.Event {
	var result = make([]models.Event, 0)
	for i := 0; i < len(chache); i++ {
		if chache[i].Time.Day() == day {
			result = append(result, chache[i])
		}
	}
	return result
}

// GetByWeek получение по неделе
func GetByWeek() []models.Event {
	var (
		result = make([]models.Event, 0)
	)
	for i := 0; i < len(chache); i++ {
		if chache[i].Time.Before(time.Now().AddDate(0, 0, 7)) && chache[i].Time.After(time.Now().AddDate(0, 0, -7)) {
			result = append(result, chache[i])
		}

	}
	return result
}

// GetByMonth - получению по месяцу
func GetByMonth(intMonth int) []models.Event {
	month := time.Month(intMonth)
	var result = make([]models.Event, 0)
	for i := 0; i < len(chache); i++ {
		if chache[i].Time.Month() == month {
			result = append(result, chache[i])
		}
	}
	return result
}

// Reset - очистка кеша для тестов
func Reset() {
	chache = make([]models.Event, 0)
}
