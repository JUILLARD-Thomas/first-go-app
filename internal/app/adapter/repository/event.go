package repository

import (
	"first-go-app/internal/app/adapter/postgresql"
	"first-go-app/internal/app/adapter/postgresql/model"

	"first-go-app/internal/app/domain"
)

// Event is the repository of domain.Event
type Event struct{}

// Get gets event
func (r Event) Get(from, to string) []domain.Event {
	db := postgresql.Connection()
	var events []model.Event
	var result postgresql.DB

	if from == "" && to == "" {
		result = db.Find(&events)
	} else if from == "" {
		result = db.Where("ts <= ?", to).Find(&events)
	} else if to == "" {
		result = db.Where("ts >= ?", from).Find(&events)
	} else {
		result = db.Where("ts BETWEEN ? AND ?", from, to).Find(&events)
	}

	if result.Error != nil {
		panic(result.Error)
	}
	result_events := make([]domain.Event, result.RowsAffected)

	for i, event := range events {
		result_events[i] = domain.Event{
			Type_enum:  event.Type_enum,
			User_agent: event.User_agent,
			Ip:         event.Ip,
			Ts:         event.Ts,
		}
	}
	return result_events
}
