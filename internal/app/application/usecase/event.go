package usecase

// NOTICE: This usecase DON'T depend on Adapter layer
import (
	"first-go-app/internal/app/domain"
	"first-go-app/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Event(r repository.IEvent) domain.Event {
	return r.Get()
}
