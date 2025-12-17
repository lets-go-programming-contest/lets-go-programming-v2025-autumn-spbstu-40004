package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jambii1/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var errExpected = errors.New("expected error")

func getMockDBRows(t *testing.T, names []string) *sqlmock.Rows {
	t.Helper()

	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}

	return rows
}

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	testData := [][]string{
		{"name1, name2"},
		{"", ""},
	}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	for rowIdx, row := range testData {
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(getMockDBRows(t, row))

		names, err := dbService.GetNames()

		require.NoError(t, err, "row: %d, error must be nil", rowIdx)
		require.Equal(t, row, names, "row: %d, expected names: %s, actual names: %s", rowIdx, row, names)
	}
}

func TestGetNamesDBQueryErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(getMockDBRows(t, nil)).
		WillReturnError(errExpected)

	names, err := dbService.GetNames()

	require.ErrorIs(t, err, errExpected, "expected error: %w, actual error: %w", errExpected, err)
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

	dbService := db.New(mockDB)

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

	dbService := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("name").
				RowError(0, errExpected),
		)

	names, err := dbService.GetNames()

	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "rows error")
}

func TestGetUniqueNamesSuccess(t *testing.T) {
	t.Parallel()

	testData := [][]string{
		{"name1, name2"},
		{"", ""},
	}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	for rowIdx, row := range testData {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(getMockDBRows(t, row))

		names, err := dbService.GetUniqueNames()

		require.NoError(t, err, "row: %d, error must be nil", rowIdx)
		require.Equal(t, row, names, "row: %d, expected names: %s, actual names: %s", rowIdx, row, names)
	}
}

func TestGetUniqueNamesDBQueryErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(getMockDBRows(t, nil)).
		WillReturnError(errExpected)

	names, err := dbService.GetUniqueNames()

	require.ErrorIs(t, err, errExpected, "expected error: %w, actual error: %w", errExpected, err)
	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "db query")
}

func TestGetUniqueNamesRowsScanningErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow(nil),
		)

	names, err := dbService.GetUniqueNames()

	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "rows scanning")
}

func TestGetUniqueNamesRowsErr(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when creating db connection", err)
	}
	defer mockDB.Close()

	dbService := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("name").
				RowError(0, errExpected),
		)

	names, err := dbService.GetUniqueNames()

	require.Nil(t, names, "names must be nil")
	require.ErrorContains(t, err, "rows error")
}
