package usecase

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/factory"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/utils/slice"
)

type NotificationUsecase interface {
	Send() error
}

type notificationUsecase struct {
	eventRepo         repository.EventRepository
	messageRepo       repository.MessageRepository
	notifiedEventRepo repository.NotifiedEventRepository
	locationRepo      repository.LocationRepository
}

func NewNotificationUsecase(eventRepo repository.EventRepository, messageRepo repository.MessageRepository, notifiedEventRepo repository.NotifiedEventRepository, locationRepo repository.LocationRepository,
) NotificationUsecase {
	return &notificationUsecase{
		eventRepo:         eventRepo,
		messageRepo:       messageRepo,
		notifiedEventRepo: notifiedEventRepo,
		locationRepo:      locationRepo,
	}
}

func (uc *notificationUsecase) Send() error {
	keywords := []string{"名古屋", "愛知"}

	events, err := uc.eventRepo.GetByKeyword(keywords)

	if err != nil {
		return err
	}

	if len(events) == 0 {
		return nil
	}

	eventsInAichi, err := uc.getEventsInAichi(events)

	if err != nil {
		return err
	}

	if len(eventsInAichi) == 0 {
		return nil
	}

	eventIDs := slice.Map(eventsInAichi, func(event entity.Event) int {
		return event.EventID
	})

	notifiedEventIDs, err := uc.notifiedEventRepo.FindByEventIDs(eventIDs)

	if err != nil {
		return err
	}

	notNotifiedEvents := slice.Filter(eventsInAichi, func(event entity.Event) bool {
		return !slice.Contains(notifiedEventIDs, event.EventID)
	})

	messages := slice.Map(notNotifiedEvents, func(event entity.Event) entity.Message {
		return factory.CreateMessage(event)
	})

	if err := uc.messageRepo.Send(messages); err != nil {
		return err
	}

	notNotifiedEventIDs := slice.Map(notNotifiedEvents, func(event entity.Event) int {
		return event.EventID
	})

	if err := uc.notifiedEventRepo.Save(notNotifiedEventIDs); err != nil {
		return err
	}

	return nil
}

func (uc *notificationUsecase) getEventsInAichi(events []entity.Event) ([]entity.Event, error) {
	eventsInAichi := []entity.Event{}

	for _, event := range events {
		if event.Coordinate == nil {
			continue
		}

		location, err := uc.locationRepo.SearchByCoordinate(*event.Coordinate)
		if err != nil {
			return nil, err
		}

		if location.Prefecture == "愛知県" {
			eventsInAichi = append(eventsInAichi, event)
		}
	}

	return eventsInAichi, nil
}
