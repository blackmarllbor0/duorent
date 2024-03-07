package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRestRoutes(pool *sql.DB) *gin.Engine {
	router := gin.Default()

	UserRoutes(pool, router)

	return router
}
