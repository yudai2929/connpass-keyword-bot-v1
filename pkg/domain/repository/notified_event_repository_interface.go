package repository

type NotifiedEventRepository interface {
	Save(eventIDs []int) error
	FindByEventIDs(eventIDs []int) ([]int, error)
}
