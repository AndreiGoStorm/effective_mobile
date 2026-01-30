package api

import (
	"effective_mobile/internal/app/requests"
	"effective_mobile/internal/app/responses"
	"effective_mobile/internal/logger"
	"effective_mobile/internal/repository"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FilterController struct {
	log *slog.Logger
	fr  *repository.FilterRepository
}

func NewFilterController(log *slog.Logger, r *repository.Repository) *FilterController {
	return &FilterController{
		log: log,
		fr:  r.FR,
	}
}

// GetSubscriptionsByUserAndService godoc
// @Summary		Total cost of all subscriptions for a selected period
// @Description  Calculating the total cost of all subscriptions for a selected period.
// @Tags         Filter
// @Accept       json
// @Produce      json
// @Param        service_name  query     string  false  "Service name (alphabetic characters only)"
// @Param        user_id       query     string  false  "User identifier (UUID format)"
// @Param        start_date    query     string  true   "Start date (Format: MM-YYYY, e.g., 01-2026)"
// @Param        end_date      query     string  true   "End date (Format: MM-YYYY, e.g., 12-2026)"
// @Success      200  {array}   responses.FilterResponse  "Успешный ответ"
// @Failure      404  {object}  responses.Response
// @Failure      422  {object}  responses.Response
// @Router        /api/v1/filters/total [get]
func (api *FilterController) GetSubscriptionsByUserAndService(c *fiber.Ctx) (err error) {
	f := new(requests.FilterByUserAndServiceRequest)
	if err := c.QueryParser(f); err != nil {
		api.log.Error("GetSubscriptions BodyParser", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.Error("invalid JSON"))
	}

	if err := validator.New().Struct(f); err != nil {
		validateErr := err.(validator.ValidationErrors)
		api.log.Error("GetSubscriptions validation error", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(f.ValidationFilterError(validateErr))
	}

	total, err := api.fr.FilterByUserAndService(f.GetFilter())
	if err != nil {
		api.log.Error("GetSubscriptions", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("failed to get subscriptions request"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.ConvertToFilterResponse(total))
}
