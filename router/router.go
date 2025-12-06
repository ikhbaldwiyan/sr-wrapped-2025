package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/users/:user_id", userHandler.GetUser)

	return router
}
