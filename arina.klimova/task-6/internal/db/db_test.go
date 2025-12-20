package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arinaklimova/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockDBRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func uniqueRows(names []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, name := range names {
		if !seen[name] {
			seen[name] = true
			result = append(result, name)
		}
	}
	return result
}

func TestNew(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	assert.NotNil(t, service)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames(t *testing.T) {
	t.Run("successful query", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Bob"}))

		names, err := service.GetNames()

		assert.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Bob"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(errors.New("db error"))

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db query:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows scanning:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errors.New("row error"))

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows error:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no rows", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDBRows([]string{}))

		names, err := service.GetNames()

		assert.NoError(t, err)
		assert.Empty(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDBService_GetUniqueNames(t *testing.T) {
	t.Run("successful query", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Bob", "Charlie"}))

		names, err := service.GetUniqueNames()

		assert.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Bob", "Charlie"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnError(errors.New("db error"))

		names, err := service.GetUniqueNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "db query:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows scanning:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows error", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errors.New("row error"))

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "rows error:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("duplicate names in result", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Alice", "Bob"}))

		names, err := service.GetUniqueNames()

		assert.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Alice", "Bob"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
