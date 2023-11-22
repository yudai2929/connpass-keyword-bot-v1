package repository

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"

type MessageRepository interface {
	Send(messages []entity.Message) error
}
