package repository

import (
	"context"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetUserByUserID(userId string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByUserID(userId string) (*models.User, error) {
	var user models.User

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}
	err = config.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
