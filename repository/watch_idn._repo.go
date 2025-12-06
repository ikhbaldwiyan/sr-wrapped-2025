package repository

import (
	"context"
	"time"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WatchIDNRepository interface {
	GetMostWatchIDN(userID string) ([]*models.WatchIDN, error)
	GetMostWatchedMembers(userID string) ([]*models.MostWatchedMember, error)
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

func (r *watchIDNRepository) GetMostWatchedMembers(userId string) ([]*models.MostWatchedMember, error) {
	var results []*models.MostWatchedMember
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	startDate := primitive.NewDateTimeFromTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	endDate := primitive.NewDateTimeFromTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"user": objectId,
			"description": bson.M{
				"$regex": primitive.Regex{Pattern: "^Watch IDN Live", Options: "i"},
			},
			"timestamp": bson.M{
				"$gte": startDate,
				"$lt":  endDate,
			},
		}}},
		{{Key: "$group", Value: bson.M{
			"_id":      "$description",
			"count":    bson.M{"$sum": 1},
			"firstLog": bson.M{"$first": "$$ROOT"},
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "idn_lives_history",
			"localField":   "firstLog.live_id",
			"foreignField": "live_id",
			"as":           "liveData",
		}}},
		{{Key: "$unwind", Value: bson.M{
			"path":                       "$liveData",
			"preserveNullAndEmptyArrays": true,
		}}},
		{{Key: "$project", Value: bson.M{
			"member": bson.M{
				"name":  "$liveData.name",
				"image": "$liveData.image",
			},
			"watch": "$count",
		}}},
		{{Key: "$sort", Value: bson.M{"watch": -1}}},
		{{Key: "$limit", Value: 10}},
	}

	cursor, err := config.Collection("activities-log").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
