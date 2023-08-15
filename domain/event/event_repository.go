package domain

type EventRepository interface {
	GetEventsByKeyword(keyword string) ([]Event, error)
}
