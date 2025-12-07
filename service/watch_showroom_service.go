package service

import (
	"fmt"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/repository"
)

type WatchShowroomService interface {
	GetMostWatchedShowroom(userId string) (*models.User, []*models.MostWatchedMember, error)
}

type watchShowroomService struct {
	watchShowroomRepo repository.WatchShowroomRepository
	userRepo          repository.UserRepository
}

func NewWatchShowroomService(watchShowroomRepo repository.WatchShowroomRepository, userRepo repository.UserRepository) WatchShowroomService {
	return &watchShowroomService{
		watchShowroomRepo: watchShowroomRepo,
		userRepo:          userRepo,
	}
}

func (service *watchShowroomService) GetMostWatchedShowroom(userId string) (*models.User, []*models.MostWatchedMember, error) {
	user, err := service.userRepo.GetUserByUserID(userId)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	mostWatchedSR, err := service.watchShowroomRepo.GetMostWatchedShowroom(userId)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return user, mostWatchedSR, nil
}
