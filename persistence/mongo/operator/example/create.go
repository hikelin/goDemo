package example

import (
	"context"
	"log"

	ExampleSchema "github.com/hikelin/goDemo/persistence/mongo/schema/example"

	OperationRunner "github.com/hikelin/goDemo/persistence/mongo/operator/runner"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Insert one item into example collection
func (operator ExampleOperator) Insert(exampleDoc ExampleSchema.Example) objectid.ObjectID {
	exampleDoc.ID = objectid.New()
	res, err := operator.collection.InsertOne(context.Background(), exampleDoc)
	if err != nil {
		log.Fatal(err)
	}

	id := res.InsertedID.(*bson.Element).Value().ObjectID()

	return id
}

// InsertMany items into example collection
func (operator ExampleOperator) InsertMany(docs []ExampleSchema.Example, jobNumber int) []interface{} {
	docsNor := convert(docs)

	ids := OperationRunner.RunOperation(docsNor, operator.InsertManyOpe, jobNumber)

	return ids
}

// InsertManyOpe items into example collection
func (operator ExampleOperator) InsertManyOpe(docs []interface{}) []interface{} {
	result, err := operator.collection.InsertMany(context.Background(), docs)
	if err != nil {
		log.Fatal(err)
	}
	ids := result.InsertedIDs
	return ids
}

func convert(docs []ExampleSchema.Example) []interface{} {
	s := make([]interface{}, len(docs))
	for i, v := range docs {
		v.ID = objectid.New()
		s[i] = v
	}
	return s
}
