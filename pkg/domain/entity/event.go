package entity

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"

type Event struct {
	EventID  int
	Title    string
	EventURL string
	*valueobject.Coordinate
}

func NewEvent(eventID int, title string, eventURL string, coordinate *valueobject.Coordinate) Event {
	return Event{
		EventID:    eventID,
		Title:      title,
		EventURL:   eventURL,
		Coordinate: coordinate,
	}
}
