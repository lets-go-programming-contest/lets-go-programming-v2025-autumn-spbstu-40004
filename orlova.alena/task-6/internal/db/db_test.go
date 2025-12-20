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

func (t *testDatabase) Query(string, ...any) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

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

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			rows []string
			want []string
		}{
			{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
			{[]string{"Single"}, []string{"Single"}},
			{[]string{}, []string{}},
			{[]string{"Test"}, []string{"Test"}},
		}

		for _, tc := range cases {
			t.Run("", func(t *testing.T) {
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

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			mock   func(sqlmock.Sqlmock)
			errMsg string
		}{
			{
				func(m sqlmock.Sqlmock) {
					m.ExpectQuery("SELECT name FROM users").WillReturnError(errMock)
				},
				"db query",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
					m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow("ok").AddRow(456)
					m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow("test").RowError(0, errMock)
					m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
				},
				"rows error",
			},
		}

		for _, tc := range cases {
			t.Run("", func(t *testing.T) {
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

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			rows []string
			want []string
		}{
			{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
			{[]string{"A", "A", "B"}, []string{"A", "B"}},
			{[]string{"Same", "Same", "Same"}, []string{"Same"}},
			{[]string{"One"}, []string{"One"}},
			{[]string{}, []string{}},
		}

		for _, tc := range cases {
			t.Run("", func(t *testing.T) {
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

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			mock   func(sqlmock.Sqlmock)
			errMsg string
		}{
			{
				func(m sqlmock.Sqlmock) {
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)
				},
				"db query",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow("valid").AddRow(nil)
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
				},
				"rows scanning",
			},
			{
				func(m sqlmock.Sqlmock) {
					rows := sqlmock.NewRows([]string{"name"}).AddRow("test").RowError(0, errMock)
					m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
				},
				"rows error",
			},
		}

		for _, tc := range cases {
			t.Run("", func(t *testing.T) {
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

	t.Run("closed connection", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		mockDB.Close()

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnError(errors.New("connection closed"))

		result, err := db.New(mockDB).GetNames()

		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("rows close", func(t *testing.T) {
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

	t.Run("large dataset", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"})
		for range 100 {
			rows.AddRow("User")
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		result, err := db.New(mockDB).GetNames()
		require.NoError(t, err)
		assert.Len(t, result, 100)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("special characters", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		names := []string{"Name-Dash", "Name_Underscore"}
		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(createRows(names...))

		result, err := db.New(mockDB).GetUniqueNames()
		require.NoError(t, err)
		assert.Equal(t, names, result)
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
