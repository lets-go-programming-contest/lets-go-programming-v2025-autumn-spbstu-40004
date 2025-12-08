package db_test

import (
	"errors"
	"github.com/15446-rus75/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	dbService := db.New(mockDB)

	tests := []struct {
		name          string
		mockNames     []string
		mockError     error
		expectedNames []string
		expectedError error
	}{
		{
			name:          "success case",
			mockNames:     []string{"Ivan", "Gena"},
			expectedNames: []string{"Ivan", "Gena"},
		},
		{
			name:          "empty result",
			mockNames:     []string{},
			expectedNames: []string{},
		},
		{
			name:          "query error",
			mockError:     errors.New("database error"),
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"name"})
			for _, name := range tc.mockNames {
				rows.AddRow(name)
			}

			mock.ExpectQuery("SELECT name FROM users").
				WillReturnRows(rows).
				WillReturnError(tc.mockError)

			names, err := dbService.GetNames()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedNames, names)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
