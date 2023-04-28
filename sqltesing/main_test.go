package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_modifyActor_PASS(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening the database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE actor").WillReturnResult(sqlmock.NewResult(1, 1))

	// act
	if err = modifyActor(db, 50, "JOE"); err != nil {
		t.Errorf("Error was not expected while modifying actors: %s", err)
	}

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

// This will fail on purpose
func Test_modifyActor_FAIL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening the database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE actor").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(fmt.Errorf("some error"))

	// act
	if err = modifyActor(db, 50, "JOE"); err != nil {
		t.Errorf("Error was not expected while modifying actors: %s", err)
	}

	// assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
