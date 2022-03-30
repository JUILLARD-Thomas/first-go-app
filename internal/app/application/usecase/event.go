package usecase

// NOTICE: This usecase DON'T depend on Adapter layer
import (
	"first-go-app/internal/app/domain"
	"first-go-app/internal/app/domain/repository"
	"time"
)

// EventArgs are arguments of Event usecase
type EventArgs struct {
	R    repository.IEvent // Interface
	FROM string
	TO   string
}

type EventArgsCreate struct {
	R          repository.IEvent // Interface
	Type_enum  string
	User_agent string
	Ip         string
}

// Event is the usecase of getting event
func Event(e EventArgs) ([]domain.Event, error) {
	return e.R.Get(e.FROM, e.TO)
}

func CreateEvent(e EventArgsCreate) (domain.Event, error) {
	event := domain.Event{
		Type_enum:  e.Type_enum,
		User_agent: e.User_agent,
		Ip:         e.Ip,
		Ts:         time.Now().Format("2006-01-02 15:04:05"),
	}
	return e.R.Post(event)
}

func CountEvent(r repository.IEvent) (int64, error) {
	return r.Count()
}
