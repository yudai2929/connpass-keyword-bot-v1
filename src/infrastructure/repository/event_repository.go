package repository

import (
	"connpass-keyword-bot-v1/domain/entity"
	"connpass-keyword-bot-v1/domain/repository"
	"connpass-keyword-bot-v1/infrastructure/response"
	"encoding/json"
	"net/http"
)

type EventRepositoryImpl struct {
	BaseURL string
	Client  *http.Client
}

func NewEventRepository(baseURL string) repository.EventRepository {
	return &EventRepositoryImpl{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

func (repo *EventRepositoryImpl) GetEventsByKeyword(keywords []string) ([]entity.Event, error) {
	keyword := convertKeywordsToString(keywords)

	url := repo.BaseURL + "/event/?keyword_or=" + keyword + "&order=3" + "&count=20"
	resp, err := repo.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := response.EventResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	events := []entity.Event{}

	for _, event := range response.Events {
		events = append(events, entity.Event{
			EventID:  event.EventID,
			Title:    event.Title,
			EventURL: event.EventURL,
		})
	}

	return events, nil
}

func convertKeywordsToString(keywords []string) string {
	keyword := ""
	for _, k := range keywords {
		keyword += k + ","
	}
	return keyword
}
