package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lolnyok/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockDatabase struct {
	queryFunc func(query string, args ...any) (*sql.Rows, error)
}

func (m *mockDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	return m.queryFunc(query, args...)
}

func TestDBService_GetNames_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	expectedNames := []string{"Alice", "Bob", "Charlie"}
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		AddRow("Charlie")

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Equal(t, expectedNames, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_QueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	testErr := errors.New("query error")
	mock.ExpectQuery("SELECT name FROM users").WillReturnError(testErr)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query")
	require.ErrorIs(t, err, testErr)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow(123)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_RowsError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	testErr := errors.New("rows error")
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, testErr)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error")
	require.ErrorIs(t, err, testErr)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_EmptyResult(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_NilRowsWithError(t *testing.T) {
	mockDB := &mockDatabase{
		queryFunc: func(query string, args ...any) (*sql.Rows, error) {
			return nil, errors.New("connection failed")
		},
	}

	service := db.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query")
}

func TestDBService_GetUniqueNames_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	expectedNames := []string{"Alice", "Bob", "Charlie"}
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		AddRow("Charlie")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Equal(t, expectedNames, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_WithDuplicates(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	expectedNames := []string{"Alice", "Bob", "Alice"}
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		AddRow("Alice")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Equal(t, expectedNames, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_QueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	testErr := errors.New("query error")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(testErr)

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query")
	require.ErrorIs(t, err, testErr)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow(123)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_RowsError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	testErr := errors.New("rows error")
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, testErr)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error")
	require.ErrorIs(t, err, testErr)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_EmptyResult(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_NilRowsWithError(t *testing.T) {
	mockDB := &mockDatabase{
		queryFunc: func(query string, args ...any) (*sql.Rows, error) {
			return nil, errors.New("connection failed")
		},
	}

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query")
}
