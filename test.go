package main

import (
	"github.com/hikelin/goDemo/persistence/mongo"
	ExampleOperator "github.com/hikelin/goDemo/persistence/mongo/operator/example"
	ExampleSchema "github.com/hikelin/goDemo/persistence/mongo/schema/example"
	"log"
	"runtime"
)

func main() {
	const (
		Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
		Ltime                         // the time in the local time zone: 01:23:23
		Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
		Llongfile                     // full file name and line number: /a/b/c/d.go:23
		Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
		LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
		LstdFlags     = Ldate | Ltime // initial values for the standard logger
	)

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

	// log.Print(len(exampleOperator.FindByQty(qtylist, 10)))

	// newEx.Status = "update"

	// log.Print(exampleOperator.UpdateByItem(newEx.Item, newEx))

	// newEx.Status = "replace"

	// log.Print(exampleOperator.ReplaceByItem(newEx.Item, newEx))

	// log.Print(exampleOperator.DeleteByItem(newEx.Item))

}
