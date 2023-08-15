package main

import (
	"connpass-keyword-bot-v1/infrastructure/repository"
	"connpass-keyword-bot-v1/interfaces/handler"
	"connpass-keyword-bot-v1/usecase"
	"net/http"
)

func main() {
	eventRepository := repository.NewEventRepository("https://connpass.com/api/v1")
	eventUsecase := usecase.NewEventUsecase(eventRepository)
	eventHandler := handler.NewEventHandler(eventUsecase)

	http.HandleFunc("/events", eventHandler.GetEventsByKeyword)
	http.ListenAndServe(":8080", nil)
}
