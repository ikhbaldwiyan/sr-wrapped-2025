package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/service"
)

type WrappedHandler struct {
	watchIDNService      service.WatchIDNService
	watchShowroomService service.WatchShowroomService
}

func NewWrappedHandler(watchIDNService service.WatchIDNService, watchShowroomService service.WatchShowroomService) *WrappedHandler {
	return &WrappedHandler{
		watchIDNService:      watchIDNService,
		watchShowroomService: watchShowroomService,
	}
}

type WrappedResponse struct {
	User     interface{} `json:"user"`
	IDN      interface{} `json:"idn"`
	Showroom interface{} `json:"showroom"`
}

func (h *WrappedHandler) GetWrappedMostWatched(c *gin.Context) {
	userId := c.Param("user_id")

	var (
		wg           sync.WaitGroup
		user         interface{}
		idnData      interface{}
		showroomData interface{}
		errIDN       error
		errShowroom  error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		var u interface{}
		u, idnData, errIDN = h.watchIDNService.GetMostWatched(userId)
		if u != nil {
			user = u
		}
	}()

	go func() {
		defer wg.Done()
		var u interface{}
		u, showroomData, errShowroom = h.watchShowroomService.GetMostWatchedShowroom(userId)
		if user == nil && u != nil {
			user = u
		}
	}()

	wg.Wait()

	if errIDN != nil {
		fmt.Printf("Error fetching IDN data: %v\n", errIDN)
	}

	if errShowroom != nil {
		fmt.Printf("Error fetching Showroom data: %v\n", errShowroom)
	}

	if errIDN != nil && errShowroom != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, WrappedResponse{
		User:     user,
		IDN:      idnData,
		Showroom: showroomData,
	})
}
