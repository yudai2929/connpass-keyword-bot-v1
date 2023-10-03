package usecase

import (
	"connpass-keyword-bot-v1/pkg/domain/entity"
	"connpass-keyword-bot-v1/pkg/domain/repository"
	"connpass-keyword-bot-v1/pkg/utils"
)

type NotificationUsecase interface {
	PostNotification() error
}

type notificationUsecase struct {
	eventRepo         repository.EventRepository
	messageRepo       repository.MessageRepository
	notifiedEventRepo repository.NotifiedEventRepository
}

func NewNotificationUsecase(eventRepo repository.EventRepository, messageRepo repository.MessageRepository, notifiedEventRepo repository.NotifiedEventRepository,
) NotificationUsecase {
	return &notificationUsecase{
		eventRepo:         eventRepo,
		messageRepo:       messageRepo,
		notifiedEventRepo: notifiedEventRepo,
	}
}

func (uc *notificationUsecase) PostNotification() error {
	keywords := []string{"名古屋", "愛知"}

	events, err := uc.eventRepo.GetByKeyword(keywords)

	if err != nil {
		return err
	}

	eventIDs := getEventIDs(events)

	notifiedEventIDs, err := uc.notifiedEventRepo.FindByEventIDs(eventIDs)

	if err != nil {
		return err
	}

	notNotifiedEvents := findNotNotifiedEvents(events, notifiedEventIDs)

	messages := createMessage(notNotifiedEvents)

	if err := uc.messageRepo.Send(messages); err != nil {
		return err
	}

	notNotifiedEventIDs := getEventIDs(notNotifiedEvents)

	if err := uc.notifiedEventRepo.Save(notNotifiedEventIDs); err != nil {
		return err
	}

	return nil
}

func getEventIDs(events []entity.Event) []int {
	eventIDs := []int{}

	for _, event := range events {
		eventIDs = append(eventIDs, event.EventID)
	}

	return eventIDs
}

func findNotNotifiedEvents(events []entity.Event, notifiedEventIDs []int) []entity.Event {
	notNotifiedEvents := []entity.Event{}

	for _, event := range events {
		if !utils.Contains(notifiedEventIDs, event.EventID) {
			notNotifiedEvents = append(notNotifiedEvents, event)
		}
	}

	return notNotifiedEvents
}

func createMessage(events []entity.Event) []entity.Message {
	messages := []entity.Message{}

	for _, event := range events {
		messages = append(messages, entity.Message{
			Text: event.Title + "\n" + event.EventURL,
		})
	}

	return messages
}
