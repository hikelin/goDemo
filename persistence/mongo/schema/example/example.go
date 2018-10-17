package example

import "github.com/mongodb/mongo-go-driver/bson/objectid"

// Example schema
type Example struct {
	ID     objectid.ObjectID `json:"id" bson:"_id"`
	Item   string            `json:"item" bson:"item"`
	Qty    int32             `json:"qty" bson:"qty"`
	Tags   []string          `json:"tags" bson:"tags"`
	Status string            `json:"status" bson:"status"`
	Size   Size              `json:"size" bson:"size"`
}

// New function is used to create a empty Example
func New() Example {
	return Example{}
}
