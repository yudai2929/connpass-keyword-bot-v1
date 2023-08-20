package repository

import (
	"connpass-keyword-bot-v1/domain/repository"

	"github.com/nedpals/supabase-go"
)

type NotifiedEventRepositoryImpl struct {
	Client *supabase.Client
}

func NewNotifiedEventRepository(supabaseURL, supabaseKey string) repository.NotifiedEventRepository {
	return &NotifiedEventRepositoryImpl{
		Client: supabase.CreateClient(supabaseURL, supabaseKey),
	}
}

type notifiedEvent struct {
	EventID int `json:"event_id"`
}

func (repo *NotifiedEventRepositoryImpl) SaveNotifiedEvents(eventIDs []int) error {
	tableName := "notified_events"

	var rows []notifiedEvent
	for _, eventID := range eventIDs {
		rows = append(rows, notifiedEvent{
			EventID: eventID,
		})
	}

	var result []notifiedEvent

	err := repo.Client.DB.From(tableName).Insert(rows).Execute(&result)

	if err != nil {
		return err
	}

	return nil
}

func (repo *NotifiedEventRepositoryImpl) FindNotifiedEventsByEventIDs(
	eventIDs []int,
) ([]int, error) {

	tableName := "notified_events"

	var result []notifiedEvent

	err := repo.Client.DB.From(tableName).Select("*").Execute(&result)

	if err != nil {
		return nil, err
	}

	var notifiedEventIDs []int

	for _, notifiedEvent := range result {
		notifiedEventIDs = append(notifiedEventIDs, notifiedEvent.EventID)
	}

	return notifiedEventIDs, nil
}
