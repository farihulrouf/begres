package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Access struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Idpagu    string             `json:"idpagu,omitempty" validate:"required"`
	UserId    string             `json:"userId,omitempty" validate:"required"`
	CreatedAt time.Time          `json:”created_at,omitempty” bson:”created_at”`
	UpdatedAt time.Time          `json:”updated_at,omitempty” bson:”updated_at”`
}
