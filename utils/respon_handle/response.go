package respon_handle

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Success response
func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  "success",
		"data":    data,
		"message": nil,
	})
}

// Error response
func Error(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"code":    code,
		"message": message,
	})
}

// ValidationError response (for binding/validation issues)
func ValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errs := make(map[string]string)
		for _, e := range ve {
			errs[e.Field()] = validationMessage(e)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Validation failed",
			"errors":  errs,
		})
	} else {
		// fallback for other kinds of errors
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
	}
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value too short"
	case "max":
		return "Value too long"
	default:
		return fe.Error()
	}
}
