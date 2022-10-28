package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pagu struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Paguopdp string             `json:"paguopdp,omitempty" validate:"required"`
	Paguorp  string             `json:"paguorp,omitempty" validate:"required"`
}
