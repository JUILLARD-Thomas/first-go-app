package repository

import "first-go-app/internal/app/domain"

// IParameter is interface of parameter repository
type IEvent interface {
	Get() domain.Event
}
