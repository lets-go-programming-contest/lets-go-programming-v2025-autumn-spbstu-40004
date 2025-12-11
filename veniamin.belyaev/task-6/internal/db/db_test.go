package db_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/belyaevEDU/task-6/internal/db"
)

var (
	getNamesRows       = []string{"Gena", "Lyoha", "Bobik", "Gena"}
	uniqueGetNamesRows = []string{"Gena", "Lyoha", "Bobik"}
)

func areStringSplicesEqual(lst, rst []string) bool {
	if len(lst) != len(rst) {
		return false
	}

	for index, elem := range lst {
		if elem != rst[index] {
			return false
		}
	}

	return true
}

func TestGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()

	defer func() {
		mock.ExpectClose()
		err := mockDB.Close()
		if err != nil {
			t.Fatalf("mockDB.Close() resulted in an error: %v", err)
		}
	}()

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	dbService := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	for _, names := range getNamesRows {
		rows = rows.AddRow(names)
	}

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	getNamesResult, err := dbService.GetNames()
	if err != nil {
		t.Fatalf("getNames error: %v", err)
	}

	if !areStringSplicesEqual(getNamesResult, getNamesRows) {
		t.Fatalf("unexpected result in getNames")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock sql expectations weren't met: %v", err)
	}
}

func TestGetUniqueNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()

	defer func() {
		mock.ExpectClose()
		err := mockDB.Close()
		if err != nil {
			t.Fatalf("mockDB.Close() resulted in an error: %v", err)
		}
	}()

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	dbService := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	for _, names := range uniqueGetNamesRows {
		rows = rows.AddRow(names)
	}

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	getNamesResult, err := dbService.GetUniqueNames()
	if err != nil {
		t.Fatalf("getNames error: %v", err)
	}

	if !areStringSplicesEqual(getNamesResult, uniqueGetNamesRows) {
		t.Fatalf("unexpected result in uniqueGetNames")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock sql expectations weren't met: %v", err)
	}
}
