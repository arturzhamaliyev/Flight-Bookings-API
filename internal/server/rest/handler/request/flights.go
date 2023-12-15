package request

import (
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

type CreateFlight struct {
	StartDate   time.Time         `json:"startDate"`
	EndDate     time.Time         `json:"endDate"`
	Departure   model.Coordinates `json:"departure"`
	Destination model.Coordinates `json:"destination"`
}
