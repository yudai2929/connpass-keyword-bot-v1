package usecase

import (
	"connpass-keyword-bot-v1/domain/event"
	"connpass-keyword-bot-v1/domain/message"
)

type NotificationUsecase interface {
	PostNotification() error
}

type notificationUsecase struct {
	eventRepo   event.EventRepository
	messageRepo message.MessageRepository
}

func NewNotificationUsecase(eventRepo event.EventRepository, messageRepo message.MessageRepository,
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
	messages := []message.Message{}

	for _, event := range events {
		messages = append(messages, message.Message{
			Text: event.Title + "\n" + event.EventURL,
		})
	}

	if err := uc.messageRepo.SendMessage(messages[:3]); err != nil {
		return err
	}

	return nil
}
