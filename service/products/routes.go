package products

import (
	"net/http"
	"strconv"

	"github.com/delapaska/1C/models"
	"github.com/delapaska/1C/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Handler struct {
	store models.ProductStore
}

func NewHandler(store models.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/:id/", h.handleGetProducts)
		api.POST("/:id/add", h.handleAddProduct)
	}
}

// handleAddProduct godoc
// @Summary Add a new product
// @Description Add a new product to an order
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param product body models.ProductPayload true "Product Payload"
// @Success 201 {string} string "Product successfully added"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 422 {object} models.ErrorResponse "Validation errors"
// @Failure 500 {object} models.ErrorResponse "Failed to add product"
// @Router /api/v1/{id}/add [post]
func (h *Handler) handleAddProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var payload models.ProductPayload
	if err := c.BindJSON(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusUnprocessableEntity, errors.Error())
		return
	}

	err = h.store.AddProduct(models.ProductPayload{
		Name:     payload.Name,
		Quantity: payload.Quantity,
		Unit:     payload.Unit,
	}, idInt)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, "Failed to add product")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": "Product successfully added"})
}

// handleGetProducts godoc
// @Summary Get products by order ID
// @Description Get all products associated with an order
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} models.ErrorResponse "Invalid request payload"
// @Failure 404 {object} models.ErrorResponse "Order not found"
// @Router /api/v1/{id}/ [get]
func (h *Handler) handleGetProducts(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Invalid order ID")
		return
	}

	products, err := h.store.GetProductsById(idInt)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, "Order not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
