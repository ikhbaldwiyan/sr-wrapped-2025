package repository

import (
	"context"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	GetUserByUserID(userID string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByUserID(userID string) (*models.User, error) {
	var user models.User
	filter := bson.M{"user_id": userID}
	err := config.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
