package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginModel struct {
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
}
