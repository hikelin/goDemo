package example

import (
	"context"
	ExampleSchema "github.com/hikelin/goDemo/persistence/mongo/schema/example"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
)

// UpdateByItem update example collection by item value matched
func (operator ExampleOperator) UpdateByItem(item string, exampleDoc ExampleSchema.Example) (int64, int64) {
	result, err := operator.collection.UpdateMany(
		context.Background(),
		bson.NewDocument(bson.EC.String("item", item)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("item", exampleDoc.Item),
				bson.EC.String("size.uom", exampleDoc.Size.Uom),
				bson.EC.String("status", exampleDoc.Status),
			),
			bson.EC.SubDocumentFromElements("$currentDate",
				bson.EC.Boolean("lastModified", true),
			),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.MatchedCount, result.ModifiedCount
}

// ReplaceByItem item in example collection by item value matched
func (operator ExampleOperator) ReplaceByItem(item string, exampleDoc ExampleSchema.Example) (int64, int64) {
	result, err := operator.collection.ReplaceOne(
		context.Background(),
		bson.NewDocument(bson.EC.String("item", item)),
		exampleDoc,
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.MatchedCount, result.ModifiedCount
}
