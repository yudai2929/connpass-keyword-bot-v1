package batch

import (
	"fmt"

	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/usecase"
)

type NotificationJob struct {
	NotificationUsecase usecase.NotificationUsecase
}

func NewNotificationJob(notificationUsecase usecase.NotificationUsecase) *NotificationJob {
	return &NotificationJob{
		NotificationUsecase: notificationUsecase,
	}
}

func (h *NotificationJob) Send() error {
	if err := h.NotificationUsecase.Send(); err != nil {
		fmt.Printf("%+v", err)
		return err
	}

	return nil
}
