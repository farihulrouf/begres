package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Totalseleksi struct {
	Tender    string  `json:"tender,omitempty" validate:"required"`
	Total     float32 `json:"total,omitempty" validate:"required"`
	Totalpagu float64 `json:"totalpagu,omitempty" validate:"required"`
	Pdn       float64 `json:"pdn,omitempty" validate:"required"`
	Idpagu    string  `json:"idpagu,omitempty" validate:"required"`
}
