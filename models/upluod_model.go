package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upload struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	//Comment string             `json:"comment,omitempty" validate:"required"`
	File      string    `json:"file,omitempty" validate:"required"`
	Idpagu    string    `json:"idpagu,omitempty" validate:"required"`
	CreatedAt time.Time `json:”created_at,omitempty” bson:”created_at”`
	UpdatedAt time.Time `json:”updated_at,omitempty” bson:”updated_at”`
}
