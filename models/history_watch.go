package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HistoryWatch struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LogName     string             `bson:"log_name" json:"log_name"`
	Description string             `bson:"description" json:"description"`
	User        primitive.ObjectID `bson:"user" json:"user"`
	LiveId      string             `bson:"live_id" json:"live_id"`
	Timestamp   primitive.DateTime `bson:"timestamp" json:"timestamp"`
}
