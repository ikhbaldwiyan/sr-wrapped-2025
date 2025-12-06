package service

import (
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/repository"
)

type WatchIDNService interface {
	GetWatchIDN(userId string) ([]*models.WatchIDN, error)
}

type watchIDNService struct {
	repo repository.WatchIDNRepository
}

func NewWatchIDNService(repo repository.WatchIDNRepository) WatchIDNService {
	return &watchIDNService{repo: repo}
}

func (s *watchIDNService) GetWatchIDN(userId string) ([]*models.WatchIDN, error) {
	return s.repo.GetMostWatchIDN(userId)
}
