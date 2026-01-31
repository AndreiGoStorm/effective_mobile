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

type SubscriptionController struct {
	log *slog.Logger
	sr  *repository.SubscriptionRepository
}

func NewSubscriptionController(log *slog.Logger, r *repository.Repository) *SubscriptionController {
	return &SubscriptionController{
		log: log,
		sr:  r.SR,
	}
}

// GetSubscriptions godoc
// @Summary		Get list of subscriptions
// @Description  Get list of subscriptions with pagination
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Param        page   query      int false "Default 1""
// @Param        size   query      int     false  "Enum (5, 10, 15, 20)"
// @Success      200  {array}   responses.SubscriptionResponse  "Успешный ответ"
// @Failure      400  {object}  responses.Response
// @Failure      422  {object}  responses.Response
// @Router        /api/v1/subscriptions [get]
func (api *SubscriptionController) GetSubscriptions(c *fiber.Ctx) (err error) {
	pr := new(requests.PaginationRequest)
	if err := c.QueryParser(pr); err != nil {
		api.log.Error("GetSubscriptions QueryParser", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.Error("invalid JSON"))
	}

	if err := validator.New().Struct(pr); err != nil {
		validateErr := err.(validator.ValidationErrors)
		api.log.Error("GetSubscriptions validation error", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(pr.ValidationPaginationError(validateErr))
	}

	total, err := api.sr.GetTotalSubscriptions()
	if err != nil {
		api.log.Error("GetSubscriptions GetTotalSubscriptions", logger.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("failed to get total subscriptions"))
	}

	subs, err := api.sr.GetSubscriptions(pr.GetSubscriptionPagination())
	if err != nil {
		api.log.Error("GetSubscriptions", logger.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("failed to get subscriptions request"))
	}

	res := responses.NewPaginatedResponse(pr.Page, pr.Size)
	return c.Status(fiber.StatusOK).JSON(res.ConvertSubscriptionsToResponse(total, subs))
}

// GetSubscriptionByID godoc
// @Summary      Show a subscription by ID
// @Description  Get subscription by ID
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Subscription ID"
// @Success      200  {object}  responses.SubscriptionResponse
// @Failure      404  {object}  responses.Response
// @Router       /api/v1/subscriptions/{id} [get]
func (api *SubscriptionController) GetSubscriptionByID(c *fiber.Ctx) (err error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		api.log.Error("invalid request id GetSubscriptionByID", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("invalid subscription id"))
	}

	sub, err := api.sr.GetSubscriptionByID(int64(id))
	if err != nil {
		api.log.Error("GetSubscriptionByID", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("failed to get request subscription"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.ConvertSubscription(sub))
}

// CreateSubscription godoc
// @Summary      Create a new subscription
// @Description  Create a new subscription
// @Param        subscription  body  requests.CreateSubscriptionRequest  true  "Subscription"
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Success      201  {object} responses.CreateSubscriptionResponse "DataResponse"
// @Failure      400  {object}  responses.Response
// @Failure      404  {object}  responses.Response
// @Failure      422  {object}  responses.Response
// @Router        /api/v1/subscriptions [post]
func (api *SubscriptionController) CreateSubscription(c *fiber.Ctx) (err error) {
	csr := new(requests.CreateSubscriptionRequest)
	if err := c.BodyParser(csr); err != nil {
		api.log.Error("CreateSubscription BodyParser", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.Error("invalid JSON"))
	}

	if err := validator.New().Struct(csr); err != nil {
		validateErr := err.(validator.ValidationErrors)
		api.log.Error("GetSubscriptions validation error", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(csr.ValidationSubscriptionError(validateErr))
	}

	ID, err := api.sr.Create(csr.GetSubscription())
	if err != nil {
		api.log.Error("Create Subscription", logger.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("failed to create request subscription"))
	}

	return c.Status(fiber.StatusCreated).JSON(responses.ConvertToCreateSubscriptionResponse(ID))
}

// UpdateSubscription godoc
// @Summary      Update a subscription
// @Description  Update a subscription
// @Param        id   path      int  true  "Subscription ID"
// @Param        subscription  body  requests.UpdateSubscriptionRequest  true  "Subscription"
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Success      202  {object} responses.SubscriptionResponse "DataResponse"
// @Failure      400  {object}  responses.Response
// @Failure      404  {object}  responses.Response
// @Failure      422  {object}  responses.Response
// @Router        /api/v1/subscriptions/{id} [put]
func (api *SubscriptionController) UpdateSubscription(c *fiber.Ctx) (err error) {
	usr := new(requests.UpdateSubscriptionRequest)
	if err := c.BodyParser(usr); err != nil {
		api.log.Error("UpdateSubscription BodyParser", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.Error("invalid JSON"))
	}

	if err := validator.New().Struct(usr); err != nil {
		validateErr := err.(validator.ValidationErrors)
		api.log.Error("GetSubscriptions validation error", logger.Err(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(usr.ValidationSubscriptionError(validateErr))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		api.log.Error("invalid request id GetSubscriptionByID", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("invalid subscription id"))
	}

	sub, err := api.sr.GetSubscriptionByID(int64(id))
	if err != nil {
		api.log.Error("GetSubscriptionByID", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("failed to update request subscription"))
	}

	err = api.sr.Update(usr.GetSubscription(sub))
	if err != nil {
		api.log.Error("Create Subscription", logger.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("failed to update request subscription"))
	}

	return c.Status(fiber.StatusAccepted).JSON(responses.ConvertSubscription(sub))
}

// DeleteSubscription godoc
// @Summary      Delete a subscription
// @Description  Delete a subscription
// @Param        id   path      int  true  "Subscription ID"
// @Tags         Subscription
// @Accept       json
// @Produce      json
// @Success      204
// @Failure      400  {object}  responses.Response
// @Failure      404  {object}  responses.Response
// @Router        /api/v1/subscriptions/{id} [delete]
func (api *SubscriptionController) DeleteSubscription(c *fiber.Ctx) (err error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		api.log.Error("invalid request id DeleteSubscription", logger.Err(err))
		return c.Status(fiber.StatusNotFound).JSON(responses.Error("invalid subscription id"))
	}

	err = api.sr.Delete(int64(id))
	if err != nil {
		api.log.Error("DeleteSubscription", logger.Err(err))
		return c.Status(fiber.StatusBadRequest).JSON(responses.Error("failed to delete subscription"))
	}

	return c.SendStatus(fiber.StatusNoContent)
}
