package model

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID        uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time

	Airplane Airplane

	Departure   Airport
	Destination Airport

	Tickets []Ticket
}

func (f Flight) AvailableTickets() []Ticket {
	availableTickets := make([]Ticket, 0)
	for _, ticket := range f.Tickets {
		if ticket.IsAvailable() {
			availableTickets = append(availableTickets, ticket)
		}
	}
	return availableTickets
}

func (f Flight) NumberOfAvailableTickets() int {
	return len(f.AvailableTickets())
}

func (f Flight) PriceOfTicketsOfEachRank() map[string]uint {
	pricesByRank := make(map[string]uint)
	for _, ticket := range f.Tickets {
		if _, ok := pricesByRank[ticket.Rank.Name.String()]; !ok {
			pricesByRank[ticket.Rank.Name.String()] = ticket.Price
		}
	}
	return pricesByRank
}

func (f Flight) NumberOfTicketsOfEachRank() map[string]int {
	numberOfTicketsByRank := make(map[string]int)
	for _, ticket := range f.Tickets {
		numberOfTicketsByRank[ticket.Rank.Name.String()]++
	}
	return numberOfTicketsByRank
}

func (f Flight) TotalNumberOfTickets() int {
	return len(f.Tickets)
}
