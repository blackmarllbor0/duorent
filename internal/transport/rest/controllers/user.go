package controllers

import (
	"context"
	"duorent.ru/internal/service"
	"duorent.ru/internal/transport/rest/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	// todo: decide something with the ctx
	users, err := uc.userService.GetAllUsers(context.Background(), uint(usersLimit))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf(err.Error()),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO
	if err := ctx.ShouldBindBodyWith(&createUserDTO, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("ctrl: user: error with unmasrshal dto: %v", err),
		})

		return
	}

	// todo: decide something with the ctx
	createdUserID, err := uc.userService.CreateUser(ctx, createUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"userId": createdUserID})
}
