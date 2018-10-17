package example

import "github.com/mongodb/mongo-go-driver/mongo"

// SchemaName is collection name in mongo
const SchemaName = "example"

// ExampleOperator structor
type ExampleOperator struct {
	collection *mongo.Collection
}

// New is the Example collection constructor
func New(db *mongo.Database) ExampleOperator {
	return ExampleOperator{collection: db.Collection(SchemaName)}
}
