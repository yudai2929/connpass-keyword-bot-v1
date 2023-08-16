package usecase

import (
	"connpass-keyword-bot-v1/domain/entity"
	"connpass-keyword-bot-v1/domain/repository"
)

type EventUsecase interface {
	GetEventsByKeyword(keyword string) ([]entity.Event, error)
}

type eventUsecase struct {
	repo repository.EventRepository
}

func NewEventUsecase(repo repository.EventRepository) EventUsecase {
	return &eventUsecase{
		repo: repo,
	}
}

func (uc *eventUsecase) GetEventsByKeyword(keyword string) ([]entity.Event, error) {
	return uc.repo.GetEventsByKeyword(keyword)
}
