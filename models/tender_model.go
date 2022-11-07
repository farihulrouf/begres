package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tender struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Paket       string             `json:"paket,omitempty" validate:"required"`
	Pagu        float32            `json:"pagu,omitempty" validate:"required"`
	Jadwal      string             `json:"jadwal,omitempty" validate:"required"`
	Pelaksanaan string             `json:"pelaksanaan,omitempty" validate:"required"`
	Pemilihan   string             `json:"pemilihan,omitempty" validate:"required"`
	Pdn         float32            `json:"pdn,omitempty" validate:"required"`
	Ket         string             `json:"ket,omitempty" validate:"required"`
	Idpagu      string             `json:"idpagu,omitempty" validate:"required"`
}
