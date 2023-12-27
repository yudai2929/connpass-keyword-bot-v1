package repository

import (
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/entity"
	"github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"
)

type LocationRepository interface {
	SearchByCoordinate(coordinate valueobject.Coordinate) (entity.Location, error)
}
