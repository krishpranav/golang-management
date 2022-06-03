package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstname" validate:"min=2,max=100"`
	LastName     *string            `json:"lastname" validate:"min=2,max=100"`
	Password     *string            `json:"password" validate:"min=6"`
	Email        *string            `json:"email" validate:"required"`
	Avatar       *string            `json:"avatar"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
	User_id      string             `json:"user_id"`
}
