package repository

import "connpass-keyword-bot-v1/pkg/domain/entity"

type EventRepository interface {
	GetEventsByKeyword(keywords []string) ([]entity.Event, error)
}
