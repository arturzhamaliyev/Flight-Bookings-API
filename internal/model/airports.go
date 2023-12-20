package model

import "github.com/google/uuid"

type Airport struct {
	ID          uuid.UUID
	Name        string
	City        string
	Country     string
	Coordinates Coordinates
}
