package repository

import (
	"context"
	"time"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WatchShowroomRepository interface {
	GetMostWatchedShowroom(userId string) ([]*models.MostWatchedMember, error)
}

type watchShowroomRepository struct{}

func NewWatchShowroomRepository() WatchShowroomRepository {
	return &watchShowroomRepository{}
}

func (r *watchShowroomRepository) GetMostWatchedShowroom(userId string) ([]*models.MostWatchedMember, error) {
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
				"$regex": primitive.Regex{Pattern: "^Watch Live", Options: "i"},
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
		{{Key: "$addFields", Value: bson.M{
			"liveIdLong": bson.M{"$toLong": "$firstLog.live_id"},
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "live_ids",
			"localField":   "liveIdLong",
			"foreignField": "live_id",
			"as":           "liveData",
		}}},
		{{Key: "$unwind", Value: bson.M{
			"path":                       "$liveData",
			"preserveNullAndEmptyArrays": true,
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "members",
			"localField":   "liveData.roomId",
			"foreignField": "room_id",
			"as":           "memberData",
		}}},
		{{Key: "$unwind", Value: bson.M{
			"path":                       "$memberData",
			"preserveNullAndEmptyArrays": true,
		}}},
		{{Key: "$project", Value: bson.M{
			"member": bson.M{
				"name":  "$memberData.stage_name",
				"image": "$memberData.image",
			},
			"watch": "$count",
		}}},
		{{Key: "$sort", Value: bson.M{"watch": -1}}},
		{{Key: "$limit", Value: 10}},
	}

	data, err := config.Collection("activities-log").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer data.Close(context.Background())

	if err = data.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
