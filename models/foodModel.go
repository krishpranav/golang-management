package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required, min=2, max=200"`
	Price     *float64           `json:"price" validate:"required"`
	FoodImage *string            `json:"food_image" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Food_id   string             `json:"food_id"`
	Menu_id   *string            `json:"menu_id" validate:"required"`
}
