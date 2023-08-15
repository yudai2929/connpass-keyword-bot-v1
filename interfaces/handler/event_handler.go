package handler

import (
	"connpass-keyword-bot-v1/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type EventHandler struct {
	EventUsecase usecase.EventUsecase
}

func NewEventHandler(eventUsecase usecase.EventUsecase) *EventHandler {
	return &EventHandler{
		EventUsecase: eventUsecase,
	}
}

func (h *EventHandler) GetEventsByKeyword(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	events, err := h.EventUsecase.GetEventsByKeyword(keyword)
	if err != nil {
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
