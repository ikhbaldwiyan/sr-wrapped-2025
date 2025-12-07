package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/service"
)

type WatchShowroomHandler struct {
	service service.WatchShowroomService
}

func NewWatchShowroomHandler(service service.WatchShowroomService) *WatchShowroomHandler {
	return &WatchShowroomHandler{service: service}
}

func (handler *WatchShowroomHandler) GetMostWatchedShowroom(context *gin.Context) {
	userId := context.Param("user_id")

	user, mostWatched, err := handler.service.GetMostWatchedShowroom(userId)
	if err != nil {
		if user == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch most watched members"})
			fmt.Println(err)
		}
		return
	}

	context.JSON(http.StatusOK, Response{
		User: user,
		Data: mostWatched,
	})
}
