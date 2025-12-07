package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Activities Log Indexes
	logModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "user", Value: 1},
				{Key: "timestamp", Value: -1},
			},
		},
		{
			Keys: bson.D{{Key: "description", Value: 1}},
		},
	}
	createIndexes(ctx, "activities-log", logModels)

	// Live IDs Indexes
	liveIdModels := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "live_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "roomId", Value: 1}},
		},
	}
	createIndexes(ctx, "live_ids", liveIdModels)

	// Members Indexes
	memberModels := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "room_id", Value: 1}},
		},
	}
	createIndexes(ctx, "members", memberModels)
}

func createIndexes(ctx context.Context, collectionName string, models []mongo.IndexModel) {
	coll := Collection(collectionName)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err := coll.Indexes().CreateMany(ctx, models, opts)
	if err != nil {
		log.Printf("Failed to create indexes for %s: %v", collectionName, err)
	} else {
		fmt.Printf("Indexes ensured for collection: %s\n", collectionName)
	}
}
