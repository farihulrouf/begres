package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Anggaran struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Paket      string             `json:"paket,omitempty" validate:"required"`
	Pagu       float64            `json:"pagu,omitempty" validate:"required"`
	Jadwal     string             `json:"jadwal,omitempty" validate:"required"`
	Pdn        float32            `json:"pdn,omitempty" validate:"required"`
	Idpagu     string             `json:"idpagu,omitempty" validate:"required"`
	UserCreate string             `json:"usercreate,omitempty" bson:”usercreate"`
	UserUpdate string             `json:"userupdate,omitempty" bson:"userupdate"`
	CreatedAt  time.Time          `json:”created_at,omitempty” bson:”created_at”`
	UpdatedAt  time.Time          `json:”updated_at,omitempty” bson:”updated_at”`
}
