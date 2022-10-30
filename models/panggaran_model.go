package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Panggaran struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Paket  string             `json:"paket,omitempty" validate:"required"`
	Pagu   string             `json:"pagu,omitempty" validate:"required"`
	Jadwal string             `json:"jadwal,omitempty" validate:"required"`
	Idpagu string             `json:"idpagu,omitempty" validate:"required"`
}
