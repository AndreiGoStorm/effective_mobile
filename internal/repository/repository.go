package repository

import (
	"effective_mobile/internal/storage"
)

type Repository struct {
	storage *storage.Storage

	SR *SubscriptionRepository
	FR *FilterRepository
}

func New(store *storage.Storage) *Repository {
	return &Repository{
		storage: store,
		SR:      NewSubscriptionRepository(store.Ctx, store.DB),
		FR:      NewFilterRepository(store.Ctx, store.DB),
	}
}
