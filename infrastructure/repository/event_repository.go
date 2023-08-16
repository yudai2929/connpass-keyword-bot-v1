package repository

import (
	"connpass-keyword-bot-v1/domain/entity"
	"connpass-keyword-bot-v1/domain/repository"
	"encoding/json"
	"fmt"
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

func (repo *EventRepositoryImpl) GetEventsByKeyword(keyword string) ([]entity.Event, error) {
	url := repo.BaseURL + "/event/?keyword_or=" + keyword + "&order=3"
	resp, err := repo.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Events []entity.Event `json:"events"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	fmt.Println(response.Events[0].Address)

	return response.Events, nil
}
