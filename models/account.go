package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAccount struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserName    string             `json:"userName" bson:"userName"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	UpVote      uint               `json:"upVote" bson:"upVote"`
	DownVote    uint               `json:"downVote" bson:"downVote"`
}

type UserAccountImportInput struct {
	UserName    string `json:"userName" bson:"userName"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
