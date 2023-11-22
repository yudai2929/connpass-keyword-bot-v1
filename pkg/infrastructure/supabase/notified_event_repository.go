package supabase

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"

	"github.com/nedpals/supabase-go"
)

type NotifiedEventRepositoryImpl struct {
	client *supabase.Client
}

func NewNotifiedEventRepository(client *supabase.Client) repository.NotifiedEventRepository {
	return &NotifiedEventRepositoryImpl{
		client,
	}
}

type notifiedEvent struct {
	EventID int `json:"event_id"`
}

func (repo *NotifiedEventRepositoryImpl) Save(eventIDs []int) error {
	tableName := "notified_events"

	var rows []notifiedEvent
	for _, eventID := range eventIDs {
		rows = append(rows, notifiedEvent{
			EventID: eventID,
		})
	}

	var result []notifiedEvent

	err := repo.client.DB.From(tableName).Insert(rows).Execute(&result)

	if err != nil {
		return err
	}

	return nil
}

func (repo *NotifiedEventRepositoryImpl) FindByEventIDs(
	eventIDs []int,
) ([]int, error) {

	tableName := "notified_events"

	var result []notifiedEvent

	err := repo.client.DB.From(tableName).Select("*").Execute(&result)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get notified events")
	}

	var notifiedEventIDs []int

	for _, notifiedEvent := range result {
		notifiedEventIDs = append(notifiedEventIDs, notifiedEvent.EventID)
	}

	return notifiedEventIDs, nil
}
