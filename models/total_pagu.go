package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Totalpagu struct {
	Totalpagu float64 `json:"totalpagu,omitempty" validate:"required"`
}
