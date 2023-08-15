package domain

import "time"

type Event struct {
	EventID       int
	Title         string
	Description   string
	EventURL      string
	StartedAt     time.Time
	EndedAt       time.Time
	Limit         int
	Place         string
	Address       string
	Latitude      string
	Longitude     string
	OwnerID       int
	OwnerNickname string
	OwnerName     string
}
