package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pagu struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Paguopdp  string             `json:"paguopdp,omitempty" validate:"required"`
	Paguorp   string             `json:"paguorp,omitempty" validate:"required"`
	Filetipe  string             `json:"filetipe,omitempty" validate:"required"`
	CreatedAt time.Time          `json:”created_at,omitempty” bson:”created_at”`
	UpdatedAt time.Time          `json:”updated_at,omitempty” bson:”updated_at”`
}
