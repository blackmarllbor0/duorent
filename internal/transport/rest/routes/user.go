package routes

import (
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/service"
	"duorent.ru/internal/transport/rest/controllers"
	"duorent.ru/internal/transport/rest/dto"
	"duorent.ru/internal/transport/rest/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	conn repository.SQLConnection,
	cfgService config.ConfigService,
	router *gin.Engine,
) {
	userRepo := postgres.NewUserRepo(conn)
	userHashRepo := postgres.NewUserHashRepo(conn)
	userHashService := service.NewUserHashService(userHashRepo, cfgService)
	userService := service.NewUserService(userRepo, userHashService)
	userController := controllers.NewUserController(userService)

	api := router.Group("/users")
	{
		api.GET("/", userController.GetAllUsers)
		api.POST("/", middleware.ValidateMiddleware(dto.CreateUserDTO{}), userController.CreateUser)
	}
}
