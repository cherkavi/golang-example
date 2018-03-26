package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func openDB(dbname, user, password string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", user, password, dbname))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return db
}

func createTable(db *sql.DB) {
	affectedrows, err := db.Query("create table test_data(id int(4) auto_increment primary key, name varchar(128) ) ")
	if err != nil {
		fmt.Println("error creating table, maybe was created previously ", err)
		return
	}
	defer affectedrows.Close()
	fmt.Println("affected rows after table creation: ", affectedrows)
}

func insertData(db *sql.DB, value string) {
	transaction, err := db.Begin()
	if err != nil {
		fmt.Println("error starting transaction ", err)
		return
	}
	defer transaction.Commit()

	if result, err := db.Exec("insert into test_data(name) values (?)", value); err != nil {
		fmt.Println("error inserting data ", err)
		return
	}
	if insertedId, err := result.LastInsertId(); err != nil {
		fmt.Printf("inserted   id: %v \n", insertedId)
	}
	if numRows, err := result.RowsAffected(); err != nil {
		fmt.Printf("affected rows: %v  \n", numRows)
	}
}

func selectData(db *sql.DB) {
	rows, err := db.Query("select * from test_data limit 5")
	if err != nil {
		fmt.Println("error selecting data ", err)
		return
	}
	defer rows.Close()

	var id string // can be int, string, float32 ...
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(" can't read data from 'test_data' ", err)
		} else {
			fmt.Printf("> %v  %v \n", id, name)
		}
	}

}

// docker pull mongo
// docker run --detach --name golang_mongo --publish 27017:27017 mongo
func main() {
	fmt.Println("---- open Mongo database --- ")
	db := openDB("technik_db", "root", "root")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println(" error during 'ping' Database ", err)
	}

	createTable(db)
	insertData(db, time.Now().Format(time.RFC822))
	selectData(db)

	fmt.Println("%v ", db)
}
