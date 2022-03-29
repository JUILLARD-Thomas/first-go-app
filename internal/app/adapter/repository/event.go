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
	var param model.Event
	result := db.First(&param, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Event{ // TODO change value
		Funds: 10,
		Btc:   8,
	}
}
