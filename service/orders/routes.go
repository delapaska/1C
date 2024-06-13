package orders

import (
	"net/http"

	"github.com/delapaska/1C/models"
	"github.com/delapaska/1C/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Handler struct {
	store models.OrderStore
}

func NewHandler(store models.OrderStore) *Handler {
	return &Handler{store: store}
}

// @Summary Создание заказа
// @Description Тут создаёшь заказ чтобы купить продуктов у Маги
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.OrderPayload true "Order Payload"
// @Success 201 {string} string "Order successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 422 {object} models.ErrorResponse "Validation errors"
// @Failure 500 {object} models.ErrorResponse "Failed to create order"
// @Router /api/v1/create [post]
func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/create", h.handleCreateOrder)
	}
}

// handleCreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input payload
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.OrderPayload true "Order Payload"
// @Success 201 {string} string "Order successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 422 {object} models.ErrorResponse "Validation errors"
// @Failure 500 {object} models.ErrorResponse "Failed to create order"
// @Router /api/v1/create [post]
func (h *Handler) handleCreateOrder(c *gin.Context) {
	var payload models.OrderPayload

	if err := c.BindJSON(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusUnprocessableEntity, errors.Error())
		return
	}

	err := h.store.CreateOrder(models.OrderPayload{
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to create order")
		return
	}

	c.JSON(http.StatusCreated, "Order successfully created")
}
