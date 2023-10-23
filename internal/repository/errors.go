package repository

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type dbError struct {
	pgErr *pgconn.PgError
}

func NewErrorsRepo() *dbError {
	return &dbError{
		pgErr: &pgconn.PgError{},
	}
}

func (db *dbError) IsUniqueViolation(err error) bool {
	return errors.As(err, &db.pgErr) && db.pgErr.Code == pgerrcode.UniqueViolation
}
