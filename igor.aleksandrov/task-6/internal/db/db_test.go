package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDBService_GetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(mock sqlmock.Sqlmock)
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Bob")
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			want:    []string{"Alice", "Bob"},
			wantErr: false,
		},
		{
			name: "query_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").WillReturnError(errors.New("db fail"))
			},
			wantErr: true,
		},
		{
			name: "scan_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "rows_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(0, errors.New("iteration fail"))
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRaw, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer dbRaw.Close()

			tt.mockFn(mock)
			service := New(dbRaw)

			got, err := service.GetNames()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDBService_GetUniqueNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		mockFn  func(mock sqlmock.Sqlmock)
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").AddRow("Charlie")
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			want:    []string{"Alice", "Charlie"},
			wantErr: false,
		},
		{
			name: "query_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errors.New("fail"))
			},
			wantErr: true,
		},
		{
			name: "scan_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "rows_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(0, errors.New("err"))
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRaw, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer dbRaw.Close()

			tt.mockFn(mock)
			service := New(dbRaw)

			got, err := service.GetUniqueNames()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
