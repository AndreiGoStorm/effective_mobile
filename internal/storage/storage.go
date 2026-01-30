package storage

import (
	"context"
	"database/sql"
	"effective_mobile/internal/config"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

type Storage struct {
	dns string
	m   *Migration

	DB  *sql.DB
	Ctx context.Context
}

func NewStorage(db config.Database) *Storage {
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.Username, db.Password, db.DBName)
	return &Storage{dns: dns}
}

func (s *Storage) Connect(ctx context.Context) error {
	db, err := sql.Open("pgx", s.dns)
	if err != nil {
		return fmt.Errorf("failed to load driver: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	s.DB = db
	s.Ctx = ctx

	s.m = NewMigration()
	if err := s.m.migrate(s.DB); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Close(_ context.Context) error {
	if err := s.DB.Close(); err != nil {
		return err
	}

	return nil
}
