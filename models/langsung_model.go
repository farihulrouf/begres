package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Langsung struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Paket  string             `json:"paket,omitempty" validate:"required"`
	Pagu   string             `json:"pagu,omitempty" validate:"required"`
	Jadwal string             `json:"jadwal,omitempty" validate:"required"`
	Pdn    string             `json:"pdn,omitempty" validate:"required"`
	Tipe   string             `json:"tipe,omitempty" validate:"required"`
	Idpagu string             `json:"idpagu,omitempty" validate:"required"`
}
