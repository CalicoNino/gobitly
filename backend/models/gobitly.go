package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gobitly struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Redirect  string             `json:"redirect,omitempty"`
	Gobitly   string             `json:"gobitly,omitempty"`
	CreatedAt string             `json:"createdAt,omitempty"`
	ExpiredAt string             `json:"expiredAt,omitempty"`
	Clicked   int                `json:"clicked"`
}
