package usecase

import domain "connpass-keyword-bot-v1/domain/event"

type EventUsecase interface {
	GetEventsByKeyword(keyword string) ([]domain.Event, error)
}

type eventUsecase struct {
	repo domain.EventRepository
}

func NewEventUsecase(repo domain.EventRepository) EventUsecase {
	return &eventUsecase{
		repo: repo,
	}
}

func (uc *eventUsecase) GetEventsByKeyword(keyword string) ([]domain.Event, error) {
	return uc.repo.GetEventsByKeyword(keyword)
}
