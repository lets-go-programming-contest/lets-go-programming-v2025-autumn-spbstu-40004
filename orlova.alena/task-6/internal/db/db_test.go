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

var errMock = errors.New("mock error")

type testDatabase struct{}

func (t *testDatabase) Query(string, ...any) (*sql.Rows, error) { return nil, nil }

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	assert.NotNil(t, service)
	assert.Equal(t, mockDB, service.DB)

	svc := db.DBService{DB: nil}
	assert.NotNil(t, svc)

	assert.NotNil(t, db.New(nil))
	assert.NotNil(t, db.New(&testDatabase{}))
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	t.Run("success cases", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name string
			rows []string
			want []string
		}{
			{"multiple", []string{"A", "B", "C"}, []string{"A", "B", "C"}},
			{"single", []string{"Single"}, []string{"Single"}},
			{"empty", []string{}, []string{}},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				defer mockDB.Close()

				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows(tc.rows...))

				result, err := db.New(mockDB).GetNames()

				require.NoError(t, err)
				assert.Equal(t, tc.want, result)
				assert.NoError(t, mock.ExpectationsWereMet())
			})
		}
	})

	t.Run("error cases", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name   string
			mock   func(sqlmock.Sqlmock)
			errMsg string
		}{
			{
				"query error",
				func(m sqlmock.Sqlmock) {
					m.ExpectQuery("SELECT name FROM users").WillReturnError(errMock)
				},
				"db query",
			},
			{
				"scan error",
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
					m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
			{
				"rows error",
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow("test").RowError(0, errMock)
					m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
				},
				"rows error",
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				defer mockDB.Close()

				tc.mock(mock)

				result, err := db.New(mockDB).GetNames()

				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, result)
				assert.NoError(t, mock.ExpectationsWereMet())
			})
		}
	})
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	t.Run("success cases", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name string
			rows []string
			want []string
		}{
			{"unique", []string{"A", "B", "C"}, []string{"A", "B", "C"}},
			{"duplicates", []string{"A", "A", "B"}, []string{"A", "B"}},
			{"single", []string{"One"}, []string{"One"}},
			{"empty", []string{}, []string{}},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				defer mockDB.Close()

				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows(tc.rows...))

				result, err := db.New(mockDB).GetUniqueNames()

				require.NoError(t, err)
				assert.Equal(t, tc.want, result)
				assert.NoError(t, mock.ExpectationsWereMet())
			})
		}
	})

	t.Run("error cases", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name   string
			mock   func(sqlmock.Sqlmock)
			errMsg string
		}{
			{
				"query error",
				func(m sqlmock.Sqlmock) {
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)
				},
				"db query",
			},
			{
				"scan null",
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				defer mockDB.Close()

				tc.mock(mock)

				result, err := db.New(mockDB).GetUniqueNames()

				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, result)
				assert.NoError(t, mock.ExpectationsWereMet())
			})
		}
	})
}

func TestEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("rows close and large data", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow("Test").CloseError(nil)
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		result, err := db.New(mockDB).GetNames()
		require.NoError(t, err)
		assert.Equal(t, []string{"Test"}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func createRows(names ...string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}

	return rows
}
