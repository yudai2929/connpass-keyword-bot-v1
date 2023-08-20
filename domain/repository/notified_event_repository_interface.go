package repository

type NotifiedEventRepository interface {
	SaveNotifiedEvents(eventIDs []int) error
	FindNotifiedEventsByEventIDs(eventIDs []int) ([]int, error)
}
