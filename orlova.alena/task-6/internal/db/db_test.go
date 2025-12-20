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
	errMock = errors.New("mock error")
	errCon  = errors.New("connection closed")
)

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

func TestDBService_Direct(t *testing.T) {
	t.Parallel()

	svc := db.DBService{DB: nil}
	assert.NotNil(t, svc)
	assert.Nil(t, svc.DB)
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
			name: "multiple names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows("Alex", "Mary", "Ivan"))
			},
			want: []string{"Alex", "Mary", "Ivan"},
		},
		{
			name: "single name",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows("Nick"))
			},
			want: []string{"Nick"},
		},
		{
			name: "empty",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows())
			},
			want: []string{},
		},
		{
			name: "one name with spaces",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").
					WillReturnRows(createRows("Jane Doe"))
			},
			want: []string{"Jane Doe"},
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
			name: "unique names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Olga", "Nick", "Helen"))
			},
			want: []string{"Olga", "Nick", "Helen"},
		},
		{
			name: "duplicates in result",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Ann", "Ann", "Victor"))
			},
			want: []string{"Ann", "Victor"},
		},
		{
			name: "all same names",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("Same", "Same", "Same"))
			},
			want: []string{"Same"},
		},
		{
			name: "single name",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows("OnlyOne"))
			},
			want: []string{"OnlyOne"},
		},
		{
			name: "empty",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnRows(createRows())
			},
			want: []string{},
		},
		{
			name: " query failed",
			mockFunc: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			errMsg:  "db query",
		},
		{
			name: "scan NULL",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows scanning",
		},
		{
			name: "scan NULL second row",
			mockFunc: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("valid").AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
			errMsg:  "rows scanning",
		},
		{
			name: "rows error",
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

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestNilDB_Panic(t *testing.T) {
	t.Parallel()

	service := db.DBService{DB: nil}

	assert.Panics(t, func() {
		_, _ = service.GetNames()
	})

	assert.Panics(t, func() {
		_, _ = service.GetUniqueNames()
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
			WillReturnError(errCon)

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
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("large dataset distinct", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		rows := sqlmock.NewRows([]string{"name"})
		for range 50 {
			rows.AddRow("User")
		}

		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

		service := db.New(mockDB)
		result, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Len(t, result, 50)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("special characters in names", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer mockDB.Close()

		names := []string{"John-Doe", "Mary_Jane", "test@email"}
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(createRows(names...))

		service := db.New(mockDB)
		result, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, names, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("rows close with error", func(t *testing.T) {
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
}

func createRows(names ...string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}

	return rows
}
