package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/service"
)

type WatchIDNHandler struct {
	service     service.WatchIDNService
	userService service.UserService
}

func NewWatchIDNHandler(service service.WatchIDNService, userService service.UserService) *WatchIDNHandler {
	return &WatchIDNHandler{service: service, userService: userService}
}

type Response struct {
	User interface{} `json:"user"`
	Data interface{} `json:"data"`
}

func (h *WatchIDNHandler) GetWatchIDN(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := h.userService.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	mostWatched, err := h.service.GetMostWatched(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch most watched members"})

		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Response{
		User: user,
		Data: mostWatched,
	})

}
