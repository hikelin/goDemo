package example

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
)

// DeleteAll will remove all items in example collection
func (operator ExampleOperator) DeleteAll() int64 {
	result, err := operator.collection.DeleteMany(
		context.Background(),
		bson.NewDocument(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}

// DeleteByItem will remove all items in example collection which matched the item value
func (operator ExampleOperator) DeleteByItem(item string) int64 {
	result, err := operator.collection.DeleteMany(
		context.Background(),
		bson.NewDocument(bson.EC.String("item", item)),
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
