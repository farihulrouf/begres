package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Link   string             `json:"link,omitempty" validate:"required"`
	Idpagu string             `json:"idpagu,omitempty" validate:"required"`
}
