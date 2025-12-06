package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/service"
)

type WatchIDNHandler struct {
	service service.WatchIDNService
}

func NewWatchIDNHandler(service service.WatchIDNService) *WatchIDNHandler {
	return &WatchIDNHandler{service: service}
}

type Response struct {
	User interface{} `json:"user"`
	Data interface{} `json:"data"`
}

func (h *WatchIDNHandler) GetWatchIDN(c *gin.Context) {
	userId := c.Param("user_id")

	user, mostWatched, err := h.service.GetMostWatched(userId)
	if err != nil {
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch most watched members"})
			fmt.Println(err)
		}
		return
	}

	c.JSON(http.StatusOK, Response{
		User: user,
		Data: mostWatched,
	})
}
