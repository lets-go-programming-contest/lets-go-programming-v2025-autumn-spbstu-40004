package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errQuery = errors.New("query failed")
	errRows  = errors.New("rows error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errQuery)

	names, err := service.GetNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "db query")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		RowError(0, errRows).
		AddRow("ignored")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errQuery)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "db query")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		RowError(0, errRows).
		AddRow("ignored")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")

	require.NoError(t, mock.ExpectationsWereMet())
}
