package db_test

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"spbstu.ru/nadia.voronina/task-6/internal/db"
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
		AddRow("Ivan").
		AddRow("Gena228")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Gena228"}, names)
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
	require.Nil(t, names)
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
}

func TestGetNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
	rows.RowError(0, errRow)

	dbService := db.New(mockDB)
	names, err := dbService.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
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
	require.Empty(t, names)
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Gena228"}, names)
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
	require.Nil(t, names)
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
	require.Nil(t, names)
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
	rows.RowError(0, errRow)

	dbService := db.New(mockDB)
	names, err := dbService.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
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
	require.Empty(t, names)
}
