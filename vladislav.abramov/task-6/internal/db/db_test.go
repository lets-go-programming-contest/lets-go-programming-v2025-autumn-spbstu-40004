package db_test

import (
	"errors"
	"testing"

	"github.com/15446-rus75/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	assert.NotNil(t, service)
	assert.Equal(t, mockDB, service.DB)
}

func TestDBService_GetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	t.Run("successful get names", func(t *testing.T) {
		expectedNames := []string{"Alice", "Bob", "Charlie"}

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			AddRow("Bob").
			AddRow("Charlie")

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, expectedNames, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty result", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"})

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Empty(t, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		expectedErr := errors.New("database connection failed")

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(expectedErr)

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Nil(t, names)
		assert.Equal(t, expectedErr, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errors.New("row error"))

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Nil(t, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow(123)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Nil(t, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("single name", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("Single")
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		names, err := service.GetNames()
		require.NoError(t, err)
		assert.Equal(t, []string{"Single"}, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("duplicate names", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("John").
			AddRow("John").
			AddRow("Jane")

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		names, err := service.GetNames()
		require.NoError(t, err)
		assert.Equal(t, []string{"John", "John", "Jane"}, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("special characters in names", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("John O'Connor").
			AddRow("Jane-Doe").
			AddRow("Иван")

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		names, err := service.GetNames()
		require.NoError(t, err)
		assert.Equal(t, []string{"John O'Connor", "Jane-Doe", "Иван"}, names)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDBService_GetNames_MultipleCalls(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows1 := sqlmock.NewRows([]string{"name"}).AddRow("First")
	rows2 := sqlmock.NewRows([]string{"name"}).AddRow("Second")

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows1)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows2)

	names1, err1 := service.GetNames()
	require.NoError(t, err1)
	assert.Equal(t, []string{"First"}, names1)

	names2, err2 := service.GetNames()
	require.NoError(t, err2)
	assert.Equal(t, []string{"Second"}, names2)

	assert.NoError(t, mock.ExpectationsWereMet())
}
