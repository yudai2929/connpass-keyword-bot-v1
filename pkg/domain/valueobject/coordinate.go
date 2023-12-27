package valueobject

import "github.com/yudai2929/connpass-keyword-bot-v1/pkg/errors"

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinate(latitude float64, longitude float64) (Coordinate, error) {
	if latitude < -90 || latitude > 90 {
		return Coordinate{}, errors.New("latitude is invalid")
	}

	if longitude < -180 || longitude > 180 {
		return Coordinate{}, errors.New("longitude is invalid")
	}

	return Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
