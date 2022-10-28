package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Panggaran struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Jumlah string             `json:"jumlah,omitempty" validate:"required"`
	Idpagu string             `json:"idpagu,omitempty" validate:"required"`
}
