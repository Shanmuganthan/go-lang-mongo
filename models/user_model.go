package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	FullName  string             `json:"name" bson:"fullName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func VerifyPassword(pwd string, hashPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)) != nil
}
