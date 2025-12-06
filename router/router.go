package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/handler"
)

func SetupRouter(userHandler *handler.UserHandler, mostWatchIdnHandler *handler.WatchIDNHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/users/:user_id", userHandler.GetUser)
	router.GET("/most-watch-idn/:user_id", mostWatchIdnHandler.GetWatchIDN)

	return router
}
