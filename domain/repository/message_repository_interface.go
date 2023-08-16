package repository

import "connpass-keyword-bot-v1/domain/entity"

type MessageRepository interface {
	SendMessage(messages []entity.Message) error
}
