package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

const (
	getAllRanksQuery = `
		SELECT * FROM ranks;
	`
)

type RanksRepository struct {
	db *sqlx.DB
}

func NewRanksRepo(db *sqlx.DB) *RanksRepository {
	return &RanksRepository{
		db: db,
	}
}

func (r *RanksRepository) GetAllRanks(ctx context.Context) ([]model.Rank, error) {
	rows, err := r.db.QueryContext(
		ctx,
		getAllRanksQuery,
	)
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrNoRows
		}
		return nil, err
	}
	defer rows.Close()

	var ranks []model.Rank
	for rows.Next() {
		var rank model.Rank
		err = rows.Scan(
			&rank.ID,
			&rank.Name,
		)
		if err != nil {
			zap.S().Info(err)
			return nil, err
		}

		ranks = append(ranks, rank)
	}

	err = rows.Err()
	if err != nil {
		zap.S().Info(err)
		return nil, err
	}

	return ranks, nil
}
