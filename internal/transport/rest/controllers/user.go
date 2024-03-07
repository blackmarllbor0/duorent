package controllers

import (
	"duorent.ru/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc UserController) GetAllUsers(ctx *gin.Context) {
	usersLimitParam := ctx.Param("limit")

	var (
		usersLimit = 0
		err        error
	)

	if strings.TrimPrefix(usersLimitParam, " ") != "" {
		usersLimit, err = strconv.Atoi(usersLimitParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Errorf("the limit must be an integer: %v", err),
			})

			return
		}
	}

	users, err := uc.userService.GetAllUsers(uint(usersLimit))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf(err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
