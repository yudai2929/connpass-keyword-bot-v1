package handler

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/usecase"
)

type NotificationHandler struct {
	NotificationUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		NotificationUsecase: notificationUsecase,
	}
}

func (h *NotificationHandler) Send() error {
	if err := h.NotificationUsecase.Send(); err != nil {
		return err
	}

	return nil
}
