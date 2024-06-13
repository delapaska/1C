package utils

import (
	"github.com/delapaska/1C/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var Validate = validator.New()

func WriteError(c *gin.Context, statusCode int, err string) {

	c.JSON(statusCode, models.ErrorResponse{Error: err})
}
