package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected")
	}
	defer db.Close()
	//make sure connect is avaliable
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	stmt, err := db.Prepare("Create Table person(id int NOT NULL AUTO_INCREMENT,first_name varchar(40),last_name varchar(40),Primary Key(id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Person Table Succesfully Migrated...")
	}

}
