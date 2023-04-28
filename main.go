package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dsn := mysql.Config{
		User:   "root",
		Passwd: "p4ssw0rd",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "sakila",
	}

	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	actorID, err := addActor("JOE", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added actor: %v\n", actorID)
}

func addActor(firstname, lastname string) (int64, error) {
	result, err := db.Exec("INSERT INTO actor (first_name, last_name) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil
}
