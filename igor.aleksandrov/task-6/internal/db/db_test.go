package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MrMels625/task-6/internal/db"
	"github.com/stretchr/testify/assert"
)

var (
	errQuery = errors.New("query failed")
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
			name: "success_multiple_names",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Ivan").
					AddRow("Oleg")
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			want:    []string{"Ivan", "Oleg"},
			wantErr: false,
		},
		{
			name: "query_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").WillReturnError(errQuery)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRaw, mock, _ := sqlmock.New()
			defer dbRaw.Close()

			tt.mockFn(mock)
			service := db.New(dbRaw)

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
			name: "success_distinct",
			mockFn: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("User1").
					AddRow("User2")
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			want:    []string{"User1", "User2"},
			wantErr: false,
		},
		{
			name: "query_error",
			mockFn: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errQuery)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRaw, mock, _ := sqlmock.New()
			defer dbRaw.Close()

			tt.mockFn(mock)
			service := db.New(dbRaw)

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
