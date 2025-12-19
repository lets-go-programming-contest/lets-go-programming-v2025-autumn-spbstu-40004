package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    service := New(db)
    
    assert.NotNil(t, service)
    assert.Equal(t, db, service.DB)
    
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    service := New(db)

    t.Run("successful query", func(t *testing.T) {
        expectedRows := sqlmock.NewRows([]string{"name"}).
            AddRow("Alice").
            AddRow("Bob")

        mock.ExpectQuery("SELECT name FROM users").
            WillReturnRows(expectedRows)

        names, err := service.GetNames()

        assert.NoError(t, err)
        assert.Equal(t, []string{"Alice", "Bob"}, names)
        require.NoError(t, mock.ExpectationsWereMet())
    })

    t.Run("query error", func(t *testing.T) {
        mock.ExpectQuery("SELECT name FROM users").
            WillReturnError(errors.New("db error"))

        names, err := service.GetNames()

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "db query:")
        assert.Nil(t, names)
        require.NoError(t, mock.ExpectationsWereMet())
    })
}

func TestDBService_GetUniqueNames(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    service := New(db)

    t.Run("successful query", func(t *testing.T) {
        expectedRows := sqlmock.NewRows([]string{"name"}).
            AddRow("Alice").
            AddRow("Bob")

        mock.ExpectQuery("SELECT DISTINCT name FROM users").
            WillReturnRows(expectedRows)

        names, err := service.GetUniqueNames()

        assert.NoError(t, err)
        assert.Equal(t, []string{"Alice", "Bob"}, names)
        require.NoError(t, mock.ExpectationsWereMet())
    })

    t.Run("query error", func(t *testing.T) {
        mock.ExpectQuery("SELECT DISTINCT name FROM users").
            WillReturnError(errors.New("db error"))

        names, err := service.GetUniqueNames()

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "db query:")
        assert.Nil(t, names)
        require.NoError(t, mock.ExpectationsWereMet())
    })
}