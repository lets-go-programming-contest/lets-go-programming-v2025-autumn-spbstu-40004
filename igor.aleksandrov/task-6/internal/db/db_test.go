package db_test

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/MrMels625/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errDBDown = errors.New("db down")
	errRow    = errors.New("row error")
)

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(errDBDown)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "db query")
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")
}

func TestGetNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
	rows.RowError(0, errRow)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")
}

func TestGetNames_EmptyResult(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Charlie")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Charlie"}, names)
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errDBDown)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "db query")
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
	rows.RowError(0, errRow)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")
}

func TestGetUniqueNames_EmptyResult(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}
