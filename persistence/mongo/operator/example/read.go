package example

import (
	"context"
	"log"

	OperationRunner "github.com/hikelin/goDemo/persistence/mongo/operator/runner"
	ExampleSchema "github.com/hikelin/goDemo/persistence/mongo/schema/example"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// FindAll retrive all items
func (operator ExampleOperator) FindAll() []ExampleSchema.Example {

	cur, err := operator.collection.Find(
		context.Background(),
		bson.NewDocument(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return convertToExample(collect(cur))
}

// FindByStatus find example by status
func (operator ExampleOperator) FindByStatus(status string) []ExampleSchema.Example {
	cursor, err := operator.collection.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("status", status)),
	)

	if err != nil {
		log.Fatal(err)
	}

	return convertToExample(collect(cursor))
}

// FindByQty find example by qty
func (operator ExampleOperator) FindByQty(qcs []int, jobNumber int) []ExampleSchema.Example {
	elist := OperationRunner.RunIntOperation(qcs, operator.FindByQtyOpe, jobNumber)

	return convertToExample(elist)
}

// FindByQtyOpe find example by qty
func (operator ExampleOperator) FindByQtyOpe(qcs []int) []interface{} {
	cursor, err := operator.collection.Find(
		context.Background(),
		bson.NewDocument(bson.EC.SubDocumentFromElements("qty",
			convertSearchInt("$in", qcs)),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	return collect(cursor)
}

// Collect the cursor data into schema
func collect(cur mongo.Cursor) []interface{} {
	var examples []interface{}

	ctx := context.Background()
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		elem := ExampleSchema.New()
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		examples = append(examples, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return examples
}

func convertToExample(elist []interface{}) []ExampleSchema.Example {
	et := []ExampleSchema.Example{}
	for _, v := range elist {
		et = append(et, v.(ExampleSchema.Example))
	}
	return et
}

func convertSearchInt(key string, qcs []int) *bson.Element {
	qns := make([]*bson.Value, len(qcs))
	for i, v := range qcs {
		qns[i] = bson.VC.Int32(int32(v))
	}
	return bson.EC.ArrayFromElements(key, qns...)
}
