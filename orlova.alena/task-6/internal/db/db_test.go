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
	dberror   = errors.New("error db")
	strerror  = errors.New("string error")
	connerror = errors.New("connection closed")
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
			name:     "empty string",
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
		tc := tc
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
			name: "error of request",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnError(dberror)
			},
			errorMsg: "db query",
		},
		{
			name: "error of scan",
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
			name: "error rows.Err",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Тест").
					RowError(0, strerror)
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(rows)
			},
			errorMsg: "rows error",
		},
	}

	for _, tc := range testCases {
		tc := tc
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
			name:     "dublicates",
			input:    []string{"Ann", "Ann", "Victor"},
			expected: []string{"Ann", "Victor"},
		},
		{
			name:     "empty string",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		tc := tc
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
			name: "error DISTINCT request",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			errorMsg: "db query",
		},
		{
			name: "NULL as a result of DISTINCT",
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
		tc := tc
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
			WillReturnError(connerror)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("large amount of requests", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		names := make([]string, 100)
		rows := sqlmock.NewRows([]string{"name"})
		for i := range 100 {
			names[i] = "User"
			rows.AddRow("User")
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Len(t, result, 100)
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
