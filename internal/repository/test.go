package repository

import (
	"context"
	"effective_mobile/internal/config"
	"effective_mobile/internal/models"
	"effective_mobile/internal/storage"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func NewTestStore() *storage.Storage {
	ctx := context.Background()
	cfg := config.LoadConfigByPath("../../config/config-testing.yaml")

	store := storage.NewStorage(cfg.Database)
	err := store.Connect(ctx)
	if err != nil {
		panic("failed to connect to database")
	}

	return store
}

func CreateSubscription() *models.Subscription {
	var sub = &models.Subscription{}
	_ = gofakeit.Struct(sub)

	sub.StartDate = gofakeit.DateRange(
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Now(),
	).Truncate(24 * time.Hour)
	sub.StartDate = time.Date(sub.StartDate.Year(), sub.StartDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	endDate := sub.StartDate.AddDate(0, 1, 0).Truncate(24 * time.Hour)
	sub.EndDate = &endDate

	return sub
}

func CreateDate(date string) time.Time {
	t, err := time.Parse("01-2006", date)
	if err != nil {
		panic("invalid period format: " + date)
	}
	return t
}
