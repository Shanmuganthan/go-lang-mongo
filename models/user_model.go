package models

import "time"

type UserModel struct {
	ID        int       `json:id`
	FullName  string    `json:name`
	Email     string    `json:email`
	Password  string    `json:password`
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
}
