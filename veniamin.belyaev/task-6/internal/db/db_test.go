package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/belyaevEDU/task-6/internal/db"
)

var (
	errQuery = errors.New("db query")
	errRows  = errors.New("rows error")
)

func closeMock(t *testing.T, mock sqlmock.Sqlmock, mockDB *sql.DB) {
	t.Helper()

	mock.ExpectClose()

	if err := mockDB.Close(); err != nil {
		t.Fatalf("mockDB.Close() resulted in an error: %v", err)
	}
}

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
	t.Parallel()

	getNamesRows := []string{"Gena", "Lyoha", "Bobik", "Gena"}

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"name"})
	for _, names := range getNamesRows {
		rows = rows.AddRow(names)
	}

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)

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

func TestGetNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(errQuery)

	dbService := db.New(mockDB)
	_, err = dbService.GetNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in query, got nil")
	}
}

func TestGetNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	_, err = dbService.GetNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in scan, got nil")
	}
}

func TestGetNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"name"}).RowError(0, errRows).AddRow("")

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	_, err = dbService.GetNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in rows, got nil")
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	uniqueGetNamesRows := []string{"Gena", "Lyoha", "Bobik"}

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

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

func TestGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errQuery)

	dbService := db.New(mockDB)
	_, err = dbService.GetUniqueNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in query, got nil")
	}
}

func TestGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	_, err = dbService.GetUniqueNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in scan, got nil")
	}
}

func TestGetUniqueNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	defer closeMock(t, mock, mockDB)

	if err != nil {
		t.Fatalf("error creating db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"name"}).RowError(0, errRows).AddRow("")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	_, err = dbService.GetUniqueNames()

	if err == nil {
		t.Fatalf("unexpected result: expected err in rows, got nil")
	}
}
