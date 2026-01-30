package repository

import (
	"effective_mobile/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubscriptionRepository(t *testing.T) {
	store := NewTestStore()

	sr := NewSubscriptionRepository(store.Ctx, store.DB)

	t.Run("get subscription by id", func(t *testing.T) {
		sub := CreateSubscription()

		ID, err := sr.Create(sub)
		require.NoError(t, err)
		require.NotEmpty(t, ID)

		actual, err := sr.GetSubscriptionByID(ID)
		require.NoError(t, err)
		require.Equal(t, ID, actual.ID)
		require.Equal(t, sub.ServiceName, actual.ServiceName)
		require.Equal(t, sub.Price, actual.Price)
		require.Equal(t, sub.UserUUID, actual.UserUUID)
		require.Equal(t, sub.StartDate, actual.StartDate)
		require.Equal(t, sub.EndDate, actual.EndDate)
	})

	t.Run("create subscription", func(t *testing.T) {
		sub := CreateSubscription()

		ID, err := sr.Create(sub)
		require.NoError(t, err)
		require.NotEmpty(t, ID)
	})

	t.Run("update subscription", func(t *testing.T) {
		ID, err := sr.Create(CreateSubscription())
		require.NoError(t, err)
		require.NotEmpty(t, ID)

		sub := CreateSubscription()
		sub.ID = ID
		err = sr.Update(sub)
		require.NoError(t, err)

		actual, err := sr.GetSubscriptionByID(ID)
		require.NoError(t, err)
		require.Equal(t, sub.ServiceName, actual.ServiceName)
		require.Equal(t, sub.Price, actual.Price)
		require.Equal(t, sub.UserUUID, actual.UserUUID)
		require.Equal(t, sub.StartDate, actual.StartDate)
		require.Equal(t, sub.EndDate, actual.EndDate)
	})

	t.Run("delete subscription", func(t *testing.T) {
		ID, err := sr.Create(CreateSubscription())
		require.NoError(t, err)
		require.NotEmpty(t, ID)

		err = sr.Delete(ID)
		require.NoError(t, err)

		actual, err := sr.GetSubscriptionByID(ID)
		require.Error(t, err)
		require.Nil(t, actual)
	})

	t.Run("list subscriptions", func(t *testing.T) {
		for i := 0; i < 6; i++ {
			sub := CreateSubscription()
			_, err := sr.Create(sub)
			require.NoError(t, err)
		}

		page1, err := sr.GetSubscriptions(&models.SubscriptionPagination{Page: 1, Size: 3})
		require.NoError(t, err)
		require.Len(t, page1, 3)

		page2, err := sr.GetSubscriptions(&models.SubscriptionPagination{Page: 2, Size: 3})
		require.NoError(t, err)
		require.Len(t, page2, 3)
	})
}
