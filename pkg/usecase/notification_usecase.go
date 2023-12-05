package usecase

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/factory"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/repository"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/utils"
)

type NotificationUsecase interface {
	Send() error
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

func (uc *notificationUsecase) Send() error {
	keywords := []string{"名古屋", "愛知"}

	events, err := uc.eventRepo.GetByKeyword(keywords)

	if err != nil {
		return err
	}

	eventIDs := utils.Map(events, func(event entity.Event) int {
		return event.EventID
	})

	notifiedEventIDs, err := uc.notifiedEventRepo.FindByEventIDs(eventIDs)

	if err != nil {
		return err
	}

	notNotifiedEvents := utils.Filter(events, func(event entity.Event) bool {
		return !utils.Contains(notifiedEventIDs, event.EventID)
	})

	messages := utils.Map(notNotifiedEvents, func(event entity.Event) entity.Message {
		return factory.CreateMessage(event)
	})

	if err := uc.messageRepo.Send(messages); err != nil {
		return err
	}

	notNotifiedEventIDs := utils.Map(notNotifiedEvents, func(event entity.Event) int {
		return event.EventID
	})

	if err := uc.notifiedEventRepo.Save(notNotifiedEventIDs); err != nil {
		return err
	}

	return nil
}
