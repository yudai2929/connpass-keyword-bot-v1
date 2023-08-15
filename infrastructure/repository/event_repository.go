package repository

import (
	domain "connpass-keyword-bot-v1/domain/event"
	"encoding/json"
	"net/http"
)

type EventRepository struct {
	BaseURL string
	Client  *http.Client
}

func NewEventRepository(baseURL string) domain.EventRepository {
	return &EventRepository{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
}

func (repo *EventRepository) GetEventsByKeyword(keyword string) ([]domain.Event, error) {
	url := repo.BaseURL + "/event/?keyword_or=" + keyword + "&order=3"
	resp, err := repo.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Events []domain.Event `json:"events"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Events, nil
}
