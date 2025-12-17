package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jambii1/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var ErrExpected = errors.New("expected error")

func getMockDbRows(t *testing.T, names []string) *sqlmock.Rows {
	t.Helper()

	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}

	return rows
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	var testData = []string{"name1, name2"}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()
	dbService := db.DBService{DB: mockDB}

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(getMockDbRows(t, testData))
	names, err := dbService.GetNames()
	require.NoError(t, err, "error must be nil")
	require.Equal(t, testData, names, "expected names: %s, actual names: %s", testData, names)
}

func TestGetNamesDBQueryErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()
	dbService := db.DBService{DB: mockDB}

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(getMockDbRows(t, nil)).
		WillReturnError(ErrExpected)
	names, err := dbService.GetNames()
	require.ErrorIs(t, err, ErrExpected, "expected error: %w, actual error: %w", ErrExpected, err)
	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "db query")
}

func TestGetNamesRowsScanningErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()
	dbService := db.DBService{DB: mockDB}

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow(nil),
		)
	names, err := dbService.GetNames()
	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "rows scanning")
}

func TestGetNamesRowsErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()
	dbService := db.DBService{DB: mockDB}

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("name").
				RowError(0, ErrExpected),
		)
	names, err := dbService.GetNames()
	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "rows error")
}
