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

func (h *WatchIDNHandler) GetWatchIDN(c *gin.Context) {
	userId := c.Param("user_id")

	watchIDN, err := h.service.GetWatchIDN(userId)

	fmt.Println(watchIDN, err)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, watchIDN)
}
