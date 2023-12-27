package entity

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"

type Location struct {
	valueobject.Coordinate
	valueobject.Address
}

func NewLocation(latitude valueobject.Coordinate, longitude valueobject.Coordinate, address valueobject.Address) Location {
	return Location{
		Coordinate: latitude,
		Address:    address,
	}
}
