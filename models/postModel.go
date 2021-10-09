package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	User        primitive.ObjectID `json:"user" bson:"user,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}

type CheckPostRequestBody struct {
	User        string `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`
}