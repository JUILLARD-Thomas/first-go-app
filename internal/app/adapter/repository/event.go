package repository

import (
	"first-go-app/internal/app/adapter/postgresql"
	"first-go-app/internal/app/adapter/postgresql/model"

	"first-go-app/internal/app/domain"
)

// Event is the repository of domain.Event
type Event struct{}

// Get event
func (r Event) Get(from, to string) ([]domain.Event, error) {
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
	result_events := make([]domain.Event, result.RowsAffected)

	if result.Error != nil {
		return result_events, result.Error
	}

	for i, event := range events {
		result_events[i] = domain.Event{
			Type_enum:  event.Type_enum,
			User_agent: event.User_agent,
			Ip:         event.Ip,
			Ts:         event.Ts,
		}
	}
	return result_events, nil
}

// Create event
func (r Event) Post(post_event domain.Event) (domain.Event, error) {
	db := postgresql.Connection()
	event := model.Event{
		Type_enum:  post_event.Type_enum,
		User_agent: post_event.User_agent,
		Ip:         post_event.Ip,
		Ts:         post_event.Ts,
	}
	result := db.Omit("id").Create(&event)

	if result.Error != nil {
		panic(result.Error)
	}

	return domain.Event{
		Type_enum:  event.Type_enum,
		User_agent: event.User_agent,
		Ip:         event.Ip,
		Ts:         event.Ts,
	}, nil
}

func (r Event) Count() (int64, error) {
	db := postgresql.Connection()
	var count int64
	var event model.Event
	var result postgresql.DB

	result = db.Find(&event).Count(&count)

	if result.Error != nil {
		panic(result.Error)
	}
	return count, nil
}
