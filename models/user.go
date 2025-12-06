package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID              string             `bson:"user_id" json:"user_id"`
	Name                string             `bson:"name" json:"name"`
	Avatar              string             `bson:"avatar" json:"avatar"`
	WatchLiveIDN        int                `bson:"watchLiveIDN" json:"watchLiveIDN"`
	WatchShowroomMember int                `bson:"watchShowroomMember" json:"watchShowroomMember"`
	TopLeaderboard      bool               `bson:"top_leaderboard" json:"top_leaderboard"`
}
