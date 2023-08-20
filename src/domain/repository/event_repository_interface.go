package repository

import "connpass-keyword-bot-v1/src/domain/entity"

type EventRepository interface {
	GetEventsByKeyword(keywords []string) ([]entity.Event, error)
}
