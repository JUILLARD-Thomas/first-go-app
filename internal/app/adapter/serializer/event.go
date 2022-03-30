package serializer

import (
	"first-go-app/internal/app/domain"
	"fmt"
)

func GroupByType(events []domain.Event) map[string][]domain.Event {
	eventByType := make(map[string][]domain.Event)

	for _, e := range events {
		if v, ok := eventByType[e.Type_enum]; ok {
			eventByType[e.Type_enum] = append(v, e)
		} else {
			eventByType[e.Type_enum] = []domain.Event{e}
		}
	}
	return eventByType
}

func GroupByOs(events []domain.Event) map[string][]domain.Event {
	eventByOs := make(map[string][]domain.Event)

	for _, e := range events {
		if v, ok := eventByOs[e.User_agent]; ok {
			eventByOs[e.User_agent] = append(v, e)
		} else {
			eventByOs[e.User_agent] = []domain.Event{e}
		}
	}
	return eventByOs
}

type TypeOs struct {
	Type, Os string
}

func (t TypeOs) String() string {
	return fmt.Sprintf("Type=%s && User-Agent=%s", t.Type, t.Os)
}

func GroupByTypeAndOs(events []domain.Event) map[string][]domain.Event {
	eventByType := make(map[string][]domain.Event)

	for _, e := range events {
		typeOs := TypeOs{e.Type_enum, e.User_agent}.String()
		if v, ok := eventByType[typeOs]; ok {
			eventByType[typeOs] = append(v, e)
		} else {
			eventByType[typeOs] = []domain.Event{e}
		}
	}
	return eventByType
}
