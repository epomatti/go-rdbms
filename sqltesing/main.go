package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var dsn = mysql.Config{
	User:   "root",
	Passwd: "p4ssw0rd",
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "sakila",
}

func main() {
	connect()
	defer db.Close()

	err := modifyActor(db, 1, "JOE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated actor")
}

func modifyActor(db *sql.DB, actorid int64, firstname string) (err error) {
	if _, err = db.Exec("UPDATE actor SET first_name = ? WHERE actor_id = ?", "JOE", 1); err != nil {
		return
	}
	return
}

func connect() {
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
