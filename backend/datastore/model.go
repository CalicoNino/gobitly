package datastore

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gobitly struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Redirect string             `json:"redirect,omitempty"`
	Gobitly  string             `json:"gobitly,omitempty"`
	Random   bool               `json:"random,omitempty"`
	Clicked  int                `json:"clicked,omitempty"`
}
