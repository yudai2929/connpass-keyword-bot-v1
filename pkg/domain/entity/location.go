package entity

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/domain/valueobject"

type Location struct {
	valueobject.Coordinate
	valueobject.Address
}

func NewLocation(coordinate valueobject.Coordinate, address valueobject.Address) Location {
	return Location{
		Coordinate: coordinate,
		Address:    address,
	}
}
