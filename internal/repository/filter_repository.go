package repository

import (
	"context"
	"database/sql"
	"effective_mobile/internal/models"
	"fmt"
)

type FilterRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewFilterRepository(ctx context.Context, db *sql.DB) *FilterRepository {
	return &FilterRepository{ctx, db}
}

func (sr *FilterRepository) FilterByUserAndService(filter *models.Filter) (sum int, err error) {
	defer sr.ctx.Done()
	const op = "FilterRepository.FilterByUserAndService"

	query := `SELECT COALESCE(SUM(price), 0) 
		FROM subscriptions 
		WHERE start_date >= $1 AND (end_date >= $2 OR end_date IS NULL)`

	args := []interface{}{filter.StartDate, filter.EndDate}
	argCount := len(args)

	if filter.ServiceName != "" {
		argCount++
		query += fmt.Sprintf(" AND service_name = $%d", argCount)
		args = append(args, filter.ServiceName)
	}

	if filter.UserUUID != "" {
		argCount++
		query += fmt.Sprintf(" AND user_uuid = $%d", argCount)
		args = append(args, filter.UserUUID)
	}

	err = sr.db.QueryRowContext(sr.ctx, query, args...).Scan(&sum)
	if err != nil {
		return 0, fmt.Errorf("%s: failed query row content: %w", op, err)
	}

	return sum, nil
}
