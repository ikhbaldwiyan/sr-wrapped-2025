package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/handler"
)

func SetupRouter(
	userHandler *handler.UserHandler,
	mostWatchIdnHandler *handler.WatchIDNHandler,
	watchShowroomHandler *handler.WatchShowroomHandler,
	wrappedHandler *handler.WrappedHandler,
) *gin.Engine {
	router := gin.Default()

	router.GET("/users/:user_id", userHandler.GetUser)
	router.GET("/most-watch-idn/:user_id", mostWatchIdnHandler.GetWatchIDN)
	router.GET("/most-watch-showroom/:user_id", watchShowroomHandler.GetMostWatchedShowroom)
	router.GET("/wrapped/:user_id", wrappedHandler.GetWrappedMostWatched)

	return router
}
