package repository

import (
	"context"
	"database/sql"
	"effective_mobile/internal/models"
	"fmt"
	"time"
)

type SubscriptionRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewSubscriptionRepository(ctx context.Context, db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{ctx, db}
}

func (sr *SubscriptionRepository) GetSubscriptionByID(ID int64) (*models.Subscription, error) {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.GetByID"

	const query = `SELECT
	    id,
		service_name,
		price,
		user_uuid,
		start_date,
		end_date,
		created_at,
		updated_at
	FROM subscriptions WHERE id = $1`

	stmp, err := sr.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("%s: prepare: %w", op, err)
	}
	defer stmp.Close()

	row := stmp.QueryRowContext(sr.ctx, ID)

	sub := &models.Subscription{}
	err = row.Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserUUID,
		&sub.StartDate,
		&sub.EndDate,
		&sub.CreatedAt,
		&sub.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: row scan: %w", op, err)
	}

	return sub, nil
}

func (sr *SubscriptionRepository) Create(sub *models.Subscription) (id int64, err error) {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.Create"

	const query = `INSERT INTO subscriptions (
		service_name,
		price,
		user_uuid,
		start_date,
		end_date
	) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	stmp, err := sr.db.PrepareContext(sr.ctx, query)
	if err != nil {
		return 0, fmt.Errorf("%s: prepare context: %w", op, err)
	}

	defer stmp.Close()

	var lastInsertID int64
	err = stmp.QueryRowContext(sr.ctx,
		sub.ServiceName,
		sub.Price,
		sub.UserUUID,
		sub.StartDate,
		sub.EndDate,
	).Scan(&lastInsertID)
	if err != nil || lastInsertID == 0 {
		return 0, fmt.Errorf("query row context: %w", err)
	}

	return lastInsertID, nil
}

func (sr *SubscriptionRepository) Update(sub *models.Subscription) error {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.Update"

	query := `UPDATE subscriptions
		SET	service_name = $1,
			price = $2,
			user_uuid = $3,
			start_date = $4,
			end_date = $5,
			updated_at = $6
		WHERE id = $7`

	stmt, err := sr.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: Prepare: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sub.ServiceName,
		sub.Price,
		sub.UserUUID,
		sub.StartDate,
		sub.EndDate,
		time.Now().Format(time.DateTime),
		sub.ID)
	if err != nil {
		return fmt.Errorf("%s: failed to load driver: %w", op, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to rows affected: %w", op, err)
	}
	if affected == 0 {
		return fmt.Errorf("%s: subscription does not exist: %w", op, err)
	}

	return nil
}

func (sr *SubscriptionRepository) Delete(ID int64) error {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.Delete"

	stmt, err := sr.db.Prepare(`delete from subscriptions where id = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(ID)
	if err != nil {
		return fmt.Errorf("%s: failed to load driver: %w", op, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to rows affected: %w", op, err)
	}
	if affected == 0 {
		return fmt.Errorf("%s: event does not exist: %w", op, err)
	}

	return nil
}

func (sr *SubscriptionRepository) GetSubscriptions(
	pag *models.SubscriptionPagination) (subs []*models.Subscription, err error) {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.GetSubscriptions"

	offset := (pag.Page - 1) * pag.Size

	const query = `SELECT 
    	id, 
    	service_name, 
    	price, 
    	user_uuid, 
    	start_date, 
    	end_date, 
    	created_at, 
    	updated_at
	FROM subscriptions
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2`

	rows, err := sr.db.QueryContext(sr.ctx, query, pag.Size, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: failed query context: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		sub := &models.Subscription{}
		err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserUUID,
			&sub.StartDate,
			&sub.EndDate,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan query: %w", op, err)
		}
		subs = append(subs, sub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subs, nil
}

func (sr *SubscriptionRepository) GetTotalSubscriptions() (total int, err error) {
	defer sr.ctx.Done()
	const op = "SubscriptionRepository.GetTotalSubscriptions"

	err = sr.db.QueryRow("SELECT COUNT(*) FROM subscriptions").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("%s: Database error (count): %w", op, err)
	}

	return total, nil
}
