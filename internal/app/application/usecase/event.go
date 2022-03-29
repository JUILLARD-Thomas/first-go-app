package usecase

// NOTICE: This usecase DON'T depend on Adapter layer
import (
	"first-go-app/internal/app/domain"
	"first-go-app/internal/app/domain/repository"
)

// Event is the usecase of getting event
func Event(r repository.IEvent) domain.Event {
	return r.Get()
}
