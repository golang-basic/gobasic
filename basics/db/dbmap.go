package main

import (
	"fmt"
	sqlite3 "github.com/mattn/go-sqlite3"
	"database/sql"
	"strings"
	"os"
	"log"
)

var DB_DRIVER string

type Person struct{
	key       sql.NullInt64
	firstName sql.NullString
	lastName  sql.NullString
	address   sql.NullString
	city      sql.NullString
}

type employee struct{
	empID       sql.NullInt64
	empName     sql.NullString
	empAge      sql.NullInt64
	empPersonId sql.NullInt64
}

func main() {
	path := getConfigPath()
	_, dir, dbname := parsePath(path)

	/**
	register the driver
	 */
	sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})
	/**
	create the database
	 */
	os.Chdir(dir)
	/*
	Where driver specifies a database driver and dataSourceName specifies database-specific connection information such as database name and authentication credentials.

Note that Open does not directly open a database connection: this is deferred until a query is made. To verify that a connection can be made before making a query, use the Ping function.
	 */
	database, err := sql.Open(DB_DRIVER, dbname)
	if err != nil {
		fmt.Println("Failed to create the handle")
	}
	if err2 := database.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}

	//Transactions are started with Begin:

	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}
	age := 27
	//The Exec, Query, QueryRow and Prepare functions already covered can be used in a transaction.

	//A transaction must end with a call to Commit or Rollback.
	/**
	Exec is used for queries where no rows are returned:
	 */
	result, err := database.Exec(
	"CREATE TABLE IF NOT EXISTS Persons ( id integer PRIMARY KEY, LastName varchar(255) NOT NULL, FirstName varchar(255), Address varchar(255), City varchar(255), CONSTRAINT uc_PersonID UNIQUE (id,LastName))",
	)
	if err != nil {
		log.Fatal(err)
	}
	result, err = database.Exec(
	"create table IF NOT EXISTS employee (employeeID integer PRIMARY KEY,name varchar(255) NOT null,age int, person_id int, FOREIGN KEY (person_id) REFERENCES persons(id), CONSTRAINT uc_empID UNIQUE (employeeID, person_id, name))",
	)
	if err != nil {
		log.Fatal(err)
	}

	/**
	 Prepare creates a prepared statement for use within a transaction.

	The returned statement operates within the transaction and can no longer be used once the transaction has been committed or rolled back.
	 */
	//	tx.Prepare("CREATE TABLE Persons ( P_Id PRIMARY KEY, LastName varchar(255) NOT NULL, FirstName varchar(255), Address varchar(255), City varchar(255), CONSTRAINT uc_PersonID UNIQUE (P_Id,LastName))")
	//	tx.Commit()

	/**
	Where result contains the last insert ID and number of rows affected. The availability of these values is dependent on the database driver.
	 */
	fmt.Println(result)


	result, err = database.Exec("Insert into Persons (id, LastName, FirstName, Address, City) values (?, ?, ?, ?, ?)", nil, "soni", "swati", "110 showers drive", "Mountain view, CA")
	if err != nil {
		log.Fatal(err)
	}
	result, err = database.Exec(
		"INSERT INTO employee (employeeID, name, age, person_id) VALUES (?, ?, ?, ?)", nil,
		"gopher",
		age, 1,
	)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()


	tx, err = database.Begin()
	if err != nil {
		log.Fatal(err)
	}
	/**
	Query is used for retrieval:
	 */

	rows, err := database.Query("SELECT * FROM employee WHERE age = ?", age)
	if err != nil {
		log.Fatal(err)
	}
	/**
	 Scan copies the columns in the current row into the values pointed at by dest.

	If an argument has type *[]byte, Scan saves in that argument a copy of the corresponding data. The copy is owned by the caller and can be modified and held indefinitely. The copy can be avoided by using an argument of type *RawBytes instead; see the documentation for RawBytes for restrictions on its use.

	If an argument has type *interface{}, Scan copies the value provided by the underlying driver without conversion. If the value is of type []byte, a copy is made and the caller owns the result.
	*/
	fmt.Println(rows.Columns())

	for rows.Next() {
		var empID sql.NullInt64
		var empName sql.NullString
		var empAge sql.NullInt64
		var empPersonId sql.NullInt64
		if err := rows.Scan(&empID, &empName, &empAge, &empPersonId); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID %d with personID:%d & name %s is age %d\n", empID.Int64, empPersonId.Int64, empName.String, empAge.Int64)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows, err = database.Query("SELECT * FROM Persons WHERE LastName = ?", "soni")
	fmt.Println(rows.Columns())

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		person := Person{}
		if err = rows.Scan(&person.key, &person.lastName, &person.firstName, &person.address, &person.city); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s lives in %s\n with key %d", person.firstName.String, person.address.String, person.key.Int64)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	database.close()
}

func parsePath(path string) (err error, databaseDirName string , dbname string) {
	pathToFile := strings.Split(path, "/")
	databaseDirName = strings.Join(pathToFile[0:len(pathToFile)-1], "/")
	return os.MkdirAll(databaseDirName, 0755), databaseDirName, pathToFile[len(pathToFile)-1]
}

func getConfigPath() (path string) {
	fmt.Println("Using sqlite3")
	fmt.Println("Enter Path to DB")
	fmt.Scanf("%s", &path)
	return
}


