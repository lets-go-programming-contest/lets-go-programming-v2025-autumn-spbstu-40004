package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/widgeiw/task-6/internal/db"
)

var (
	errDB  = errors.New("error db")
	errStr = errors.New("string error")
	errCon = errors.New("connection closed")
)

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	assert.NotNil(t, service)
	assert.Equal(t, mockDB, service.DB)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "several names",
			input:    []string{"Alex", "Mary", "Ivan"},
			expected: []string{"Alex", "Mary", "Ivan"},
		},
		{
			name:     "single name",
			input:    []string{"Nick"},
			expected: []string{"Nick"},
		},
		{
			name:     "empty result",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "names with spaces",
			input:    []string{"Jane Doe", "Samara Morgan"},
			expected: []string{"Jane Doe", "Samara Morgan"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			mock.ExpectQuery("SELECT name FROM users").
				WillReturnRows(createMockRows(tc.input))

			service := db.New(mockDB)
			result, err := service.GetNames()

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetNames_Errors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		mockSetup func(sqlmock.Sqlmock)
		errorMsg  string
	}{
		{
			name: "query error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnError(errDB)
			},
			errorMsg: "db query",
		},
		{
			name: "scan error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("okay").
					AddRow(123)
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			errorMsg: "rows scanning",
		},
		{
			name: "rows.Err error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Test").
					RowError(0, errStr)
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			errorMsg: "rows error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tc.mockSetup(mock)

			service := db.New(mockDB)
			result, err := service.GetNames()

			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errorMsg)
			assert.Nil(t, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "unique names",
			input:    []string{"Olga", "Nick", "Helen"},
			expected: []string{"Olga", "Nick", "Helen"},
		},
		{
			name:     "duplicates",
			input:    []string{"Ann", "Ann", "Victor"},
			expected: []string{"Ann", "Victor"},
		},
		{
			name:     "empty result",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			mock.ExpectQuery("SELECT DISTINCT name FROM users").
				WillReturnRows(createMockRows(tc.input))

			service := db.New(mockDB)
			result, err := service.GetUniqueNames()

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUniqueNames_Errors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		mockSetup func(sqlmock.Sqlmock)
		errorMsg  string
	}{
		{
			name: "DISTINCT query error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			errorMsg: "db query",
		},
		{
			name: "NULL in DISTINCT result",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Pavel").
					AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(rows)
			},
			errorMsg: "rows scanning",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tc.mockSetup(mock)

			service := db.New(mockDB)
			result, err := service.GetUniqueNames()

			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errorMsg)
			assert.Nil(t, result)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("nil Database in constructor", func(t *testing.T) {
		t.Parallel()

		service := db.New(nil)
		assert.NotNil(t, service)
		assert.Nil(t, service.DB)
	})

	t.Run("closed connection", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		mockDB.Close()

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(errCon)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("large amount of records", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"})

		for range 100 {
			rows.AddRow("User")
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Len(t, result, 100)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows Close error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Test").
			CloseError(errors.New("close error"))

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Test"}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetUniqueNames with many duplicates", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"})
		for i := 0; i < 50; i++ {
			rows.AddRow("Duplicate")
		}

		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Len(t, result, 50)
		assert.Equal(t, "Duplicate", result[0])
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetUniqueNames with special characters", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		names := []string{"Name-Dash", "Name_Underscore", "Name.Dot"}
		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(createMockRows(names))

		service := db.New(mockDB)
		result, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Equal(t, names, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func createMockRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})

	for _, name := range names {
		rows.AddRow(name)
	}

	return rows
}
