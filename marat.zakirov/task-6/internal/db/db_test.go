package db_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ZakirovMS/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

type rowTestDB struct {
	names       []string
	errExpected error
}

type customError string

func (e customError) Error() string {
	return string(e)
}

var (
	errNoNames      = customError("NoNames")
	errRepeatedName = customError("RepeatedName")
	errQueryFailed  = customError("query failed")
	errRowIteration = customError("row iteration error")
)

func mockDBRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}

	return rows
}

func TestGetName(t *testing.T) {
	t.Parallel()

	testTableGetName := []rowTestDB{
		{
			names: []string{"nameSt", "nameNd", "nameRd"},
		},
		{
			names: []string{"repeatingName", "repeatingName", "repeatingName"},
		},
		{
			names:       nil,
			errExpected: errNoNames,
		},
	}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range testTableGetName {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDBRows(row.names)).WillReturnError(row.errExpected)

		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)

			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	testTableGetUniqueName := []rowTestDB{
		{
			names: []string{"nameSt", "nameNd", "nameRd"},
		},
		{
			names:       []string{"repeatingName", "repeatingName", "repeatingName"},
			errExpected: errRepeatedName,
		},
		{
			names:       nil,
			errExpected: errNoNames,
		},
	}

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range testTableGetUniqueName {
		query := "SELECT DISTINCT name FROM users"
		mock.ExpectQuery(query).WillReturnRows(mockDBRows(row.names)).WillReturnError(row.errExpected)

		names, err := dbService.GetUniqueNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)

			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	defer mockDB.Close()

	require.NoError(t, err, "should create mock DB without error")

	dbService := db.New(mockDB)
	require.NotNil(t, dbService, "DBService should not be nil")
	require.NotNil(t, dbService.DB, "DB field should not be nil")
}

func TestGetNameQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}
	expectedErr := errQueryFailed
	mock.ExpectQuery("SELECT name FROM users").WillReturnError(expectedErr)

	names, err := dbService.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "db query")
	require.Nil(t, names)
}

func TestGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}
	expectedErr := errQueryFailed
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(expectedErr)

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "db query")
	require.Nil(t, names)
}

func TestGetNameScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning")
	require.Nil(t, names)
}

func TestGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning")
	require.Nil(t, names)
}

func TestGetNameRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("test1").
		AddRow("test2").
		RowError(1, errRowIteration)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error")
	require.Nil(t, names)
}

func TestGetUniqueNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("test1").
		AddRow("test2").
		RowError(1, errRowIteration)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error")
	require.Nil(t, names)
}
