package models

import "time"

type Terminal struct {
	ID           uint64 // bigint
	ClientID     string
	ClientSecret string
	UUID         string // UUID храним как строку
}

type Transaction struct {
	ID            uint64
	TerminalUUID  string // UUID тоже строкой
	OrderID       string
	Amount        float64
	Status        string
	CreatedAt     time.Time
	StatusChanged time.Time
	Code          string
	Message       string
}
