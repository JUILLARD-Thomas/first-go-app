package usecase

// NOTICE: This usecase DON'T depend on Adapter layer
import (
	"first-go-app/internal/app/domain"
	"first-go-app/internal/app/domain/repository"
)

// EventArgs are arguments of Event usecase
type EventArgs struct {
	R   repository.IEvent // Interface
	MAX string
	MIN string
}

// Event is the usecase of getting event
func Event(e EventArgs) []domain.Event {
	return e.R.Get(e.MIN, e.MAX)
}
