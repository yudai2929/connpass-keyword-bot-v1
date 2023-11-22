package repository

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"

type EventRepository interface {
	GetByKeyword(keywords []string) ([]entity.Event, error)
}
