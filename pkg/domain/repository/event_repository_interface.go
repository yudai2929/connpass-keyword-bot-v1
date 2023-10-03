package repository

import "connpass-keyword-bot-v1/pkg/domain/entity"

type EventRepository interface {
	GetByKeyword(keywords []string) ([]entity.Event, error)
}
