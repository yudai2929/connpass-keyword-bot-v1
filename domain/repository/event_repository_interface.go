package repository

import "connpass-keyword-bot-v1/domain/entity"

type EventRepository interface {
	GetEventsByKeyword(keywords []string) ([]entity.Event, error)
}
