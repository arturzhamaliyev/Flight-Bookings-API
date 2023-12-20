package repository

import (
	"context"
	"fmt"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	insertTicketsQuery = `
		INSERT INTO tickets(
			id,
			flight_id,
			rank_id,
			price,
			created_at,
			updated_at
		)
		VALUES
	`
)

type TicketsRepository struct {
	db *sqlx.DB
}

func NewTicketsRepo(db *sqlx.DB) *TicketsRepository {
	return &TicketsRepository{
		db: db,
	}
}

func (t *TicketsRepository) InsertTickets(ctx context.Context, tickets []model.Ticket) error {
	insertTicketsQueryInstance := insertTicketsQuery
	vals := []interface{}{}
	counter := 1

	for _, ticket := range tickets {
		insertTicketsQueryInstance += fmt.Sprintf(`($%d, $%d, $%d, $%d, $%d, $%d),`,
			counter,
			counter+1,
			counter+2,
			counter+3,
			counter+4,
			counter+5,
		)
		counter += 6

		vals = append(vals,
			ticket.ID,
			ticket.Flight.ID,
			ticket.Rank.ID,
			ticket.Price,
			ticket.CreatedAt,
			ticket.CreatedAt,
		)
	}
	insertTicketsQueryInstance = insertTicketsQueryInstance[:len(insertTicketsQueryInstance)-1] + `;`

	st, err := t.db.PrepareContext(
		ctx,
		insertTicketsQueryInstance,
	)
	if err != nil {
		zap.S().Info(err)
		return err
	}

	_, err = st.ExecContext(
		ctx,
		vals...,
	)

	// _, err := t.db.ExecContext(
	// 	ctx,
	// 	insertTicketsQueryInstance,
	// 	vals...,
	// )
	if err != nil {
		zap.S().Info(err)
		return err
	}
	return nil
}
