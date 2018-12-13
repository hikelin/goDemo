package test

import (
	"demo/persistence/mysql"
	"demo/persistence/mysql/models"
)

// TestMysql test mysql
func TestMysql() {
	db := mysql.Connect()
	defer db.Close()

	// stmtIns, err := db.Prepare("INSERT INTO square_num VALUES( ?, ? )") // ? = placeholder
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	// stmtOut, err := db.Prepare("SELECT id FROM square_num WHERE number = ?")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtOut.Close()

	tx := mysql.BeginTx(db)
	// Insert square numbers for 0-24 in the database
	// for i := 0; i < 25; i++ {
	// 	_, err = tx.Stmt(stmtIns).Exec(i, (i * i)) // Insert tuples (i, i^2)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	// }

	// mysql.Rollback(tx)
	newExmple := models.Example{Item: "sds", Qty: 32, Tags: "xx", Status: "active", Size: "xx"}

	operator := mysql.CreateOperator(db, tx)

	operator.Insert(newExmple)

	mysql.Commit(tx)

}
