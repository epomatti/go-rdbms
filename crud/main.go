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

	// Connect and ping the database
	connect()
	defer db.Close()

	ping()

	// Add an actor
	actorID, err := addActor("JOE", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added actor: %v\n", actorID)

	// Get an actor
	actors, err := getActor(actorID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor found: %v\n", actors)

	// Update
	rowsUpdated, err := updateActor("JAMES", actorID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total actors updated: %d\n", rowsUpdated)

	// Delete
	rowsDeleted, err := deleteActor(actorID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total actors deleted: %d\n", rowsDeleted)

}

func connect() {
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}

func ping() {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
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

func getActor(actorID int64) ([]Actor, error) {
	var actors []Actor

	result, err := db.Query("SELECT actor_id, first_name, last_name FROM actor WHERE actor_id = ?", actorID)

	if err != nil {
		return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
	}

	defer result.Close()

	for result.Next() {
		var actor Actor
		if err := result.Scan(&actor.actor_id, &actor.first_name, &actor.last_name); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
		actors = append(actors, actor)

		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetActor %v: %v", actorID, err)
		}
	}
	return actors, nil
}

func updateActor(firstname string, actorid int64) (int64, error) {
	result, err := db.Exec("UPDATE actor SET first_name = ? WHERE actor_id = ?", firstname, actorid)
	if err != nil {
		return 0, fmt.Errorf("updateActor: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateActor: %v", err)
	}
	return id, nil
}

func deleteActor(actorid int64) (int64, error) {
	result, err := db.Exec("DELETE from actor WHERE actor_id = ?", actorid)
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteActor: %v", err)
	}
	return id, nil
}
