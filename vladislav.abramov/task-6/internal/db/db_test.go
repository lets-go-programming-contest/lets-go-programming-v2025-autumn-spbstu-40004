package db_test

import (
	"errors"
	"testing"

	"github.com/15446-rus75/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
)

var (
	errQuery = errors.New("query failed")
	errRows  = errors.New("rows error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(names) != 2 || names[0] != "Alice" || names[1] != "Bob" {
		t.Fatalf("unexpected names: %v", names)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errQuery)

	_, err := service.GetNames()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetNames()
	if err == nil {
		t.Fatalf("expected scan error, got nil")
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(names) != 2 {
		t.Fatalf("unexpected result: %v", names)
	}
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errQuery)

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatalf("expected scan error, got nil")
	}
}

func TestGetNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		RowError(0, errRows).
		AddRow("ignored")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetNames()
	if err == nil {
		t.Fatalf("expected rows error, got nil")
	}
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		RowError(0, errRows).
		AddRow("ignored")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatalf("expected rows error, got nil")
	}
}
