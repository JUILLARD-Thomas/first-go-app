package repository

import "first-go-app/internal/app/domain"

// IEvent is interface of event repository
type IEvent interface {
	Get(string, string) []domain.Event
	Post(domain.Event) (domain.Event, error)
	Count() (int64, error)
}
