package routes

import (
	"duorent.ru/internal/repository"
	"github.com/gin-gonic/gin"
)

func InitRestRoutes(conn repository.SQLConnection) *gin.Engine {
	router := gin.Default()

	UserRoutes(conn, router)

	return router
}
