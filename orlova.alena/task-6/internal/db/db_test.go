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

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	assert.NotNil(t, service)
	assert.Equal(t, mockDB, service.DB)

	service2 := db.New(nil)
	assert.NotNil(t, service2)
	assert.Nil(t, service2.DB)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mockFunc func(sqlmock.Sqlmock)
		want     []string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "success - multiple names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows("Alex", "Mary", "Ivan"))
			},
			want: []string{"Alex", "Mary", "Ivan"},
		},
		{
			name: "success - single name",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows("Nick"))
			},
			want: []string{"Nick"},
		},
		{
			name: "success - empty",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows())
			},
			want: []string{},
		},
		{
			name: "error - query failed",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnError(errMock)
			},
			wantErr: true,
			errMsg:  "db query",
		},
		{
			name: "error - scan failed first row",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows scanning",
		},
		{
			name: "error - scan failed second row",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("ok").AddRow(456)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows scanning",
		},
		{
			name: "error - rows error",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("test").RowError(0, errMock)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tc.mockFunc(mock)

			service := db.New(mockDB)
			got, err := service.GetNames()

			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		mockFunc func(sqlmock.Sqlmock)
		want     []string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "success - unique names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Olga", "Nick", "Helen"))
			},
			want: []string{"Olga", "Nick", "Helen"},
		},
		{
			name: "success - duplicates in result",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Ann", "Ann", "Victor"))
			},
			want: []string{"Ann", "Victor"},
		},
		{
			name: "success - all same names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Same", "Same", "Same"))
			},
			want: []string{"Same"},
		},
		{
			name: "success - single name",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("OnlyOne"))
			},
			want: []string{"OnlyOne"},
		},
		{
			name: "success - empty",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows())
			},
			want: []string{},
		},
		{
			name: "error - query failed",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			errMsg:  "db query",
		},
		{
			name: "error - scan NULL",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows scanning",
		},
		{
			name: "error - rows error",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("test").RowError(0, errMock)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tc.mockFunc(mock)

			service := db.New(mockDB)
			got, err := service.GetUniqueNames()

			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
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

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.Error(t, err)
		assert.Nil(t, result)
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

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Len(t, result, 100)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows close coverage", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Single").CloseError(nil)

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Single"}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

// Вспомогательная функция
func createRows(names ...string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}

	return rows
}
