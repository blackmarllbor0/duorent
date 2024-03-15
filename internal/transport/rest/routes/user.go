package routes

import (
	"duorent.ru/internal/repository"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/service"
	"duorent.ru/internal/transport/rest/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(pool repository.SQLConnection, router *gin.Engine) {
	userRepo := postgres.NewUserRepo(pool)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := router.Group("/users")
	{
		api.GET("/", userController.GetAllUsers)
	}
}
