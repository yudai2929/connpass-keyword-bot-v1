package repository

import "connpass-keyword-bot-v1/pkg/domain/entity"

type MessageRepository interface {
	Send(messages []entity.Message) error
}
