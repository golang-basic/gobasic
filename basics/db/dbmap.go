package main

import (
	"fmt"
	sqlite3 "github.com/mattn/go-sqlite3"
	"database/sql"
)

var DB_DRIVER string

func main() {
	fmt.Println("Using sqlite3")
	fmt.Println("Enter Path to DB")
	var path string
	fmt.Scanf("%s", &path)

	sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})

//	db, err := go-sqlite3
	database, err := sql.Open(DB_DRIVER, "newDataSource")
	if err != nil{
		fmt.Println("failed to create the handle")
	}
	if err2 := database.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}

}


