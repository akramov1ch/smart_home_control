package models

type Device struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Type string `bson:"type"`
}
