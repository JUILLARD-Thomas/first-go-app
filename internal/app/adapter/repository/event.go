package repository

import (
	"first-go-app/internal/app/adapter/postgresql"
	"first-go-app/internal/app/adapter/postgresql/model"

	"first-go-app/internal/app/domain"
)

// Parameter is the repository of domain.Parameter
type Event struct{}

// Get gets parameter
func (r Event) Get() domain.Event {
	db := postgresql.Connection()
	var event model.Event
	result := db.Take(&event)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Event{
		Type_enum:  event.Type_enum,
		User_agent: event.User_agent,
		Ip:         event.Ip,
		Ts:         event.Ts,
	}
}
