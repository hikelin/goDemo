package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect is Shared connection
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3301)/go_demo")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// BeginTx will apply the operation
func BeginTx(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	return tx
}

// Rollback will restore the uncommit operation
func Rollback(tx *sql.Tx) {
	err := tx.Rollback()

	if err != nil {
		log.Fatal(err)
	}
}

func Commit(tx *sql.Tx) {
	err := tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

}
