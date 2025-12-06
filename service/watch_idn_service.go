package service

import (
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/repository"
)

type WatchIDNService interface {
	GetWatchIDN(userId string) ([]*models.HistoryWatch, error)
	GetMostWatched(userId string) (*models.User, []*models.MostWatchedMember, error)
}

type watchIDNService struct {
	repo     repository.WatchIDNRepository
	userRepo repository.UserRepository
}

func NewWatchIDNService(repo repository.WatchIDNRepository, userRepo repository.UserRepository) WatchIDNService {
	return &watchIDNService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *watchIDNService) GetWatchIDN(userId string) ([]*models.HistoryWatch, error) {
	return s.repo.GetMostWatchIDN(userId)
}

func (s *watchIDNService) GetMostWatched(userId string) (*models.User, []*models.MostWatchedMember, error) {
	user, err := s.userRepo.GetUserByUserID(userId)
	if err != nil {
		return nil, nil, err
	}

	mostWatched, err := s.repo.GetMostWatchedMembers(userId)
	if err != nil {
		return nil, nil, err
	}

	return user, mostWatched, nil
}
