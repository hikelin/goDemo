package test

import (
	"demo/persistence/mongo"
	ExampleOperator "demo/persistence/mongo/operator/example"
	ExampleSchema "demo/persistence/mongo/schema/example"
	"log"
	"runtime"
)

// TestMongo will call mongo feature
func TestMongo() {

	db := mongo.Connect()

	exampleOperator := ExampleOperator.New(db)

	tags := []string{"xx"}
	newEx := ExampleSchema.Example{Item: "test", Qty: 19, Tags: tags, Status: "D"}

	var newExs []ExampleSchema.Example
	var qtylist []int

	for i := 0; i < 100000; i++ {
		newEx.Qty = int32(i)
		newExs = append(newExs, newEx)
	}

	for i := 0; i < 5000; i++ {
		qtylist = append(qtylist, i)
	}

	runtime.GOMAXPROCS(4)
	ids := exampleOperator.InsertMany(newExs, 10)

	log.Print(len(ids))

	log.Print(len(exampleOperator.FindByQty(qtylist, 10)))

	// newEx.Status = "update"

	// log.Print(exampleOperator.UpdateByItem(newEx.Item, newEx))

	// newEx.Status = "replace"

	// log.Print(exampleOperator.ReplaceByItem(newEx.Item, newEx))

	// log.Print(exampleOperator.DeleteByItem(newEx.Item))
}
