package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	v "github.com/go-playground/validator/v10"
	"net/http"
)

var validator = v.New()

func ValidateMiddleware[T comparable](dto T) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check that the req body obj is present
		if ctx.Request.ContentLength == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			ctx.Abort()
			return
		}

		if err := ctx.ShouldBindBodyWith(&dto, binding.JSON); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if err := validator.Struct(dto); err != nil {
			msgs := make(map[string]string, len(err.(v.ValidationErrors)))
			for _, err := range err.(v.ValidationErrors) {
				fmt.Println(err.ActualTag())

				msgs[err.Field()] = err.Tag()
			}

			ctx.JSON(http.StatusBadRequest, gin.H{"error": msgs})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
