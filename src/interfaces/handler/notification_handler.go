package handler

import (
	"connpass-keyword-bot-v1/src/usecase"
	"fmt"
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
		fmt.Println(err)
		return err
	}
	return nil
}
