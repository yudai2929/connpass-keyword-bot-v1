package external

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"
)

type EventRepositoryImpl struct {
	BaseURL string
	Client  *http.Client
}

func NewEventRepository() repository.EventRepository {
	return &EventRepositoryImpl{
		BaseURL: config.Env.ConnpassURL,
		Client:  http.DefaultClient,
	}
}

func (repo *EventRepositoryImpl) GetByKeyword(keywords []string) ([]entity.Event, error) {
	keyword := convertKeywordsToString(keywords)

	url := repo.BaseURL + "/event/?keyword_or=" + keyword + "&order=3" + "&count=20"
	resp, err := repo.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := EventResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	return convertEvent(response)
}

func convertKeywordsToString(keywords []string) string {
	keyword := ""
	for _, k := range keywords {
		keyword += k + ","
	}
	return keyword
}

func convertEvent(response EventResponse) ([]entity.Event, error) {
	events := []entity.Event{}

	for _, event := range response.Events {
		if event.Lat == "" || event.Lon == "" {
			continue
		}

		lat, err := strconv.ParseFloat(event.Lat, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse float")
		}

		lon, err := strconv.ParseFloat(event.Lon, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse float")
		}

		coordinate, err := valueobject.NewCoordinate(lat, lon)
		if err != nil {
			return nil, err
		}

		events = append(events, entity.Event{
			EventID:    event.EventID,
			Title:      event.Title,
			EventURL:   event.EventURL,
			Coordinate: &coordinate,
		})
	}

	return events, nil
}

type EventResponse struct {
	ResultsReturned  int `json:"results_returned"`
	ResultsStart     int `json:"results_start"`
	ResultsAvailable int `json:"results_available"`
	Events           []struct {
		EventID     int    `json:"event_id"`
		Title       string `json:"title"`
		Catch       string `json:"catch"`
		Description string `json:"description"`
		EventURL    string `json:"event_url"`
		HashTag     string `json:"hash_tag"`
		StartedAt   string `json:"started_at"`
		EndedAt     string `json:"ended_at"`
		Limit       int    `json:"limit"`
		EventType   string `json:"event_type"`
		Series      struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"series"`
		Address          string `json:"address"`
		Place            string `json:"place"`
		Lat              string `json:"lat"`
		Lon              string `json:"lon"`
		OwnerID          int    `json:"owner_id"`
		OwnerNickname    string `json:"owner_nickname"`
		OwnerDisplayName string `json:"owner_display_name"`
		Accepted         int    `json:"accepted"`
		Waiting          int    `json:"waiting"`
		UpdatedAt        string `json:"updated_at"`
	} `json:"events"`
}
