package main

import (
	"context"
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

	ctx := context.Background()
	actorID, err := txActor(ctx, "JOEN", "BERRY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor ID: %v\n", actorID)
}

func txActor(ctx context.Context, firstname, lastname string) (int64, error) {

	// Being the transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("adding actor failed: %v", err)
	}

	// Defer a rollback in case of a failure
	defer tx.Rollback()

	// Check if name exists
	var actID int64
	selectQuery := "SELECT actor_id from actor where first_name = ? and last_name = ?"
	if err = tx.QueryRowContext(ctx, selectQuery, firstname, lastname).Scan(&actID); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(fmt.Errorf("actor does not exist: %w", err))
		} else {
			return 0, fmt.Errorf("txActor: %v", err)
		}
	}

	// Rollback if actor exists
	if actID > 0 {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("txActor: %v", err)
		}
		fmt.Println("Actor already exist: ", actID)
		fmt.Println("*** Transaction rolling back ***")
		return actID, nil
	}

	// Create a new row
	insertStatement := "INSERT INTO actor (first_name, last_name) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, insertStatement, firstname, lastname)
	if err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	}

	// Get the ID of the actor just inserted
	NewActorID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("txActor: %v", err)
	} else {
		fmt.Println("New actor created: ", NewActorID)
	}

	return NewActorID, nil
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
