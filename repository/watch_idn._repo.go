package repository

import (
	"context"
	"time"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WatchIDNRepository interface {
	GetMostWatchIDN(userID string) ([]*models.WatchIDN, error)
}

type watchIDNRepository struct{}

func NewWatchIDNRepository() WatchIDNRepository {
	return &watchIDNRepository{}
}

func (r *watchIDNRepository) GetMostWatchIDN(userId string) ([]*models.WatchIDN, error) {
	var watchIDNs []*models.WatchIDN
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	startDate := primitive.NewDateTimeFromTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	endDate := primitive.NewDateTimeFromTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))

	filter := bson.M{
		"user":     objectId,
		"log_name": "Watch",
		"description": bson.M{
			"$regex": "IDN Live",
		},
		"timestamp": bson.M{
			"$gte": startDate,
			"$lt":  endDate,
		},
	}
	opts := options.Find().SetLimit(10).SetSort(bson.D{{Key: "timestamp", Value: -1}})

	cursor, err := config.Collection("activities-log").Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var watchIDN models.WatchIDN
		if err := cursor.Decode(&watchIDN); err != nil {
			return nil, err
		}
		watchIDNs = append(watchIDNs, &watchIDN)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return watchIDNs, nil
}
