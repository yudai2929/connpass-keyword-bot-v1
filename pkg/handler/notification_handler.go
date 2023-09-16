package handler

import (
	"connpass-keyword-bot-v1/pkg/usecase"
)

type NotificationHandler struct {
	NotificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		NotificationUsecase: notificationUsecase,
	}
}

func (h *NotificationHandler) PostNotification() error {
	err := h.NotificationUsecase.PostNotification()
	if err != nil {
		return err
	}
	return nil
}
