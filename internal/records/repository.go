package records

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type RecordsRepository interface {
	GetAllRecordsByDate(ctx context.Context, startDate, endDate string) ([]RecordModel, error)
}

type recordsRepository struct {
	db *sqlx.DB
}

func NewRecordsRepository(db *sqlx.DB) RecordsRepository {
	return &recordsRepository{
		db: db,
	}
}

func (r *recordsRepository) GetAllRecordsByDate(ctx context.Context, startDate, endDate string) ([]RecordModel, error) {
	Records := []RecordModel{}
	query := "SELECT id, name, marks, createdAt FROM records WHERE createdAt > $1 AND createdAt < $2"
	err := r.db.SelectContext(ctx, &Records, query, startDate, endDate)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return Records, nil
}
