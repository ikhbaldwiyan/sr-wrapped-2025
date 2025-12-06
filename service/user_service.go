package service

import (
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/repository"
)

type UserService interface {
	GetUser(userID string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(userID string) (*models.User, error) {
	return s.repo.GetUserByUserID(userID)
}
