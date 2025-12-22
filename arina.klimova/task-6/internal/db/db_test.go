package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arinaklimova/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errDB  = errors.New("db error")
	errRow = errors.New("row error")
)

func mockDBRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})

	for _, name := range names {
		rows = rows.AddRow(name)
	}

	return rows
}

func TestNew(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	assert.NotNil(t, service)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames(t *testing.T) {
	t.Parallel()

	t.Run("successful query", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Bob"}))

		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Bob"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(errDB)

		names, err := service.GetNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "db query:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "rows scanning:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errRow)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "rows error:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no rows", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDBRows([]string{}))

		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Empty(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDBService_GetUniqueNames(t *testing.T) {
	t.Parallel()

	t.Run("successful query", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Bob", "Charlie"}))

		names, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Bob", "Charlie"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("query error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnError(errDB)

		names, err := service.GetUniqueNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "db query:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("scan error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "rows scanning:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errRow)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "rows error:")
		assert.Nil(t, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("duplicate names in result", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		service := db.New(sqlDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDBRows([]string{"Alice", "Alice", "Bob"}))

		names, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Alice", "Alice", "Bob"}, names)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
