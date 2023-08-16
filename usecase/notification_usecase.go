package usecase

import (
	"connpass-keyword-bot-v1/domain/entity"
	"connpass-keyword-bot-v1/domain/repository"
)

type NotificationUsecase interface {
	PostNotification() error
}

type notificationUsecase struct {
	eventRepo   repository.EventRepository
	messageRepo repository.MessageRepository
}

func NewNotificationUsecase(eventRepo repository.EventRepository, messageRepo repository.MessageRepository,
) NotificationUsecase {
	return &notificationUsecase{
		eventRepo:   eventRepo,
		messageRepo: messageRepo,
	}
}

func (uc *notificationUsecase) PostNotification() error {
	events, err := uc.eventRepo.GetEventsByKeyword("名古屋")

	if err != nil {
		return err
	}
	messages := []entity.Message{}

	for _, event := range events {
		messages = append(messages, entity.Message{
			Text: event.Title + "\n" + event.EventURL,
		})
	}

	// if err := uc.messageRepo.SendMessage(messages[:3]); err != nil {
	// 	return err
	// }

	return nil
}
