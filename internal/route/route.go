package route

import (
	"effective_mobile/cmd/application/api"
	"effective_mobile/internal/repository"
	"effective_mobile/internal/storage"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	subscriptionController *api.SubscriptionController
	filterController       *api.FilterController
}

func New(store *storage.Storage, log *slog.Logger) *Route {
	repo := repository.New(store)
	return &Route{
		subscriptionController: api.NewSubscriptionController(log, repo),
		filterController:       api.NewFilterController(log, repo),
	}
}

func V1Routes(app *fiber.App, apiHandlers *Route) *fiber.Router {
	v1 := app.Group("/api/v1")

	SubscriptionRoutes(v1, apiHandlers)
	FilterRoutes(v1, apiHandlers)

	return &v1
}

func SubscriptionRoutes(router fiber.Router, apiHandlers *Route) {
	routerSubscriptions := router.Group("/subscriptions")
	{
		routerSubscriptions.Get("", apiHandlers.subscriptionController.GetSubscriptions)
		routerSubscriptions.Get("/:id<int>", apiHandlers.subscriptionController.GetSubscriptionByID)
		routerSubscriptions.Post("", apiHandlers.subscriptionController.CreateSubscription)
		routerSubscriptions.Put("/:id", apiHandlers.subscriptionController.UpdateSubscription)
		routerSubscriptions.Delete("/:id<int>", apiHandlers.subscriptionController.DeleteSubscription)
	}
}

func FilterRoutes(router fiber.Router, apiHandlers *Route) {
	routerFilters := router.Group("/filters")
	{
		routerFilters.Get("/total", apiHandlers.filterController.GetSubscriptionsByUserAndService)
	}
}
