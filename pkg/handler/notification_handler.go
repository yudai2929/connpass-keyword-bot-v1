package handler

import (
	"connpass-keyword-bot-v1/pkg/usecase"
	"context"
	"log"

	"github.com/cloudevents/sdk-go/v2/event"
)

type NotificationHandler struct {
	NotificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		NotificationUsecase: notificationUsecase,
	}
}

func (h *NotificationHandler) PostNotification(ctx context.Context, e event.Event) error {
	log.Printf("START Event Context: %+v\n", e.Context)
	if err := h.NotificationUsecase.PostNotification(); err != nil {
		return err
	}
	log.Printf("END Event Context: %+v\n", e.Context)
	return nil
}
