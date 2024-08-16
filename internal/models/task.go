package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoTask struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	CreatedAt   string             `bson:"created_at"`
}

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
