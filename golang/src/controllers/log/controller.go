package logcontroller

import (
	"fmt"
	"logger/internal/redis"
	responses "logger/src/controllers/base"
	logmodel "logger/src/models/log_model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ReportError(c *gin.Context) {
	var errors []string

	var logRequest logmodel.ReportableLog

	c.BindJSON(&logRequest)

	validate := validator.New()

	err := validate.Struct(logRequest)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Validation error:", err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}

		c.JSON(http.StatusOK, responses.ValidationErrorResponse(errors))
		return
	}

	res := redis.Lpush("logs", logRequest)

	c.JSON(http.StatusOK, responses.CreatedResponse(res))
}
