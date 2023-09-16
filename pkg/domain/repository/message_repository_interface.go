package repository

import "connpass-keyword-bot-v1/pkg/domain/entity"

type MessageRepository interface {
	SendMessage(messages []entity.Message) error
}
