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
		setup    func(sqlmock.Sqlmock)
		expected []string
		wantErr  bool
	}{
		{
			name: "multiple names",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alice").
					AddRow("Bob")
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Alice", "Bob"},
		},
		{
			name: "single name",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Single")
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Single"},
		},
		{
			name: "empty result",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"})
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expected: []string{},
		},
		{
			name: "query error",
			setup: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT name FROM users").WillReturnError(errMock)
			},
			wantErr: true,
		},
		{
			name: "scan error",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "rows error",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Test").RowError(0, errMock)
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tt.setup(mock)

			service := db.New(mockDB)
			result, err := service.GetNames()

			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(sqlmock.Sqlmock)
		expected []string
		wantErr  bool
	}{
		{
			name: "unique names",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alice").
					AddRow("Bob")
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Alice", "Bob"},
		},
		{
			name: "duplicates",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alice").
					AddRow("Alice")
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Alice", "Alice"},
		},
		{
			name: "empty result",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"})
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expected: []string{},
		},
		{
			name: "query error",
			setup: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
		{
			name: "scan error",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "rows error",
			setup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Test").RowError(0, errMock)
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer mockDB.Close()

			tt.setup(mock)

			service := db.New(mockDB)
			result, err := service.GetUniqueNames()

			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
