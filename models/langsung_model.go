package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Langsung struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Paket       string             `json:"paket,omitempty" validate:"required"`
	Pagu        float64            `json:"pagu,omitempty" validate:"required"`
	Jadwal      string             `json:"jadwal,omitempty" validate:"required"`
	Pdn         float32            `json:"pdn,omitempty,truncate" validate:"required"`
	Tipe        string             `json:"tipe,omitempty" validate:"required"`
	Ket         string             `json:"ket,omitempty" validate:"required"`
	Pelaksanaan string             `json:"pelaksanaan,omitempty" validate:"required"`
	Pemilihan   string             `json:"pemilihan,omitempty" validate:"required"`
	Tender      string             `json:"tender,omitempty" validate:"required"`
	Idpagu      string             `json:"idpagu,omitempty" validate:"required"`
	CreatedAt   time.Time          `json:”created_at,omitempty” bson:”created_at”`
	UpdatedAt   time.Time          `json:”updated_at,omitempty” bson:”updated_at”`
}
