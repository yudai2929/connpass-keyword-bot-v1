package repository

import "connpass-keyword-bot-v1/src/domain/entity"

type MessageRepository interface {
	SendMessage(messages []entity.Message) error
}
