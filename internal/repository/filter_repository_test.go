package repository

import (
	"effective_mobile/internal/models"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestFilterRepository(t *testing.T) {
	store := NewTestStore()

	sr := NewSubscriptionRepository(store.Ctx, store.DB)
	fr := NewFilterRepository(store.Ctx, store.DB)

	t.Run("filter subscriptions by userUUID", func(t *testing.T) {
		userUUID := gofakeit.UUID()

		var total int
		for i := 0; i < 5; i++ {
			sub := CreateSubscription()

			sub.UserUUID = userUUID
			sub.StartDate = CreateDate("05-2025")
			sub.EndDate = nil

			_, err := sr.Create(sub)
			require.NoError(t, err)

			total += sub.Price
		}

		filter := &models.Filter{
			UserUUID:  userUUID,
			StartDate: CreateDate("04-2025"),
		}

		actual, err := fr.FilterByUserAndService(filter)
		require.NoError(t, err)
		require.Equal(t, total, actual)
	})

	t.Run("filter subscriptions by userUUID and serviceName", func(t *testing.T) {
		userUUID := gofakeit.UUID()
		serviceName := gofakeit.AppName()

		var total int
		for i := 0; i < 3; i++ {
			sub := CreateSubscription()

			sub.ServiceName = serviceName
			sub.UserUUID = userUUID
			sub.StartDate = CreateDate("07-2025")
			sub.EndDate = nil
			_, err := sr.Create(sub)
			require.NoError(t, err)

			total += sub.Price
		}

		filter := &models.Filter{
			ServiceName: serviceName,
			UserUUID:    userUUID,
			StartDate:   CreateDate("05-2025"),
		}

		actual, err := fr.FilterByUserAndService(filter)
		require.NoError(t, err)
		require.Equal(t, total, actual)
	})

	t.Run("filter subscriptions by dates", func(t *testing.T) {
		userUUID := gofakeit.UUID()

		var total int
		for i := 0; i < 2; i++ {
			sub := CreateSubscription()

			sub.UserUUID = userUUID
			sub.StartDate = CreateDate("05-2025")
			endDate := sub.StartDate.AddDate(0, 1, 0)
			sub.EndDate = &endDate

			_, err := sr.Create(sub)
			require.NoError(t, err)

			total += sub.Price
		}

		filter := &models.Filter{
			UserUUID:  userUUID,
			StartDate: CreateDate("04-2025"),
		}

		endDate := filter.StartDate.AddDate(0, 1, 0)
		filter.EndDate = &endDate

		actual, err := fr.FilterByUserAndService(filter)
		require.NoError(t, err)
		require.Equal(t, total, actual)
	})
}
