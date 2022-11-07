package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Tendertotalpaket struct {
	Paket     string  `json:"paket,omitempty" validate:"required"`
	Total     float32 `json:"total,omitempty" validate:"required"`
	Totalpagu float32 `json:"totalpagu,omitempty" validate:"required"`
	Pdn       float32 `json:"pdn,omitempty" validate:"required"`
}
