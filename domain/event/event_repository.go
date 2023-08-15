package event

type EventRepository interface {
	GetEventsByKeyword(keyword string) ([]Event, error)
}
