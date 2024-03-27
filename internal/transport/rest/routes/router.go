package routes

import (
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository"
	"github.com/gin-gonic/gin"
)

func InitRestRoutes(conn repository.SQLConnection, cfgService config.ConfigService) *gin.Engine {
	router := gin.Default()

	UserRoutes(conn, cfgService, router)

	return router
}
