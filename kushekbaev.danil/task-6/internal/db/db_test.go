package db_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	Mdb "github.com/Z-1337/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

func TestNewDBServiceStoresDB(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectationsWereMet()

	service := Mdb.New(db)
	require.Equal(t, db, service.DB, "expected DB to be set")
}

func TestGetNames_ReturnsAllNames(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Garen").
		AddRow("Orianna").
		AddRow("Leblanc").
		AddRow("Azir").
		AddRow("Neeko").
		AddRow("Gangplank")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.NoError(t, err)

	require.Len(t, names, 6, "expected 6 names")
	require.Equal(t, "Garen", names[0], "first name should be Garen")
	require.Equal(t, "Orianna", names[1], "second name should be Orianna")
	require.Equal(t, "Leblanc", names[2], "third name should be Leblanc")
	require.Equal(t, "Azir", names[3], "fourth name should be Azir")
	require.Equal(t, "Neeko", names[4], "fifth name should be Neeko")
	require.Equal(t, "Gangplank", names[5], "sixth name should be Gangplank")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_KeepDuplicates(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Cho`Gath").
		AddRow("Pantheon").
		AddRow("Pantheon")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.NoError(t, err)

	require.Len(t, names, 3, "expected 3 names with duplicates")
	require.Equal(t, []string{"Cho`Gath", "Pantheon", "Pantheon"}, names)

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_EmptyResult(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.NoError(t, err)
	require.Empty(t, names, "expected empty slice")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_QueryErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(sql.ErrConnDone)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "db query", "error should contain 'db query'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_ScanErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Singed").
		AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "rows scanning", "error should contain 'rows scanning'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_RowsErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"})
	rows.AddRow("Zaahen")
	rows.RowError(0, sql.ErrTxDone)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "rows error", "error should contain 'rows error'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_FullSetReturned(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Teemo").
		AddRow("Yone").
		AddRow("Pyke").
		AddRow("Ezreal").
		AddRow("Samira").
		AddRow("Corki").
		AddRow("Ryze").
		AddRow("Irelia").
		AddRow("Poppy").
		AddRow("Twitch")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Len(t, names, 10, "expected 10 names (with duplicates from DISTINCT)")

	expected := []string{"Teemo", "Yone", "Pyke", "Ezreal", "Samira", "Corki", "Ryze", "Irelia", "Poppy", "Twitch"}
	for i, name := range names {
		require.Equal(t, expected[i], name, "name mismatch at index %d", i)
	}

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_OneName(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Jinx").
		AddRow("Jinx").
		AddRow("Jinx")
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Len(t, names, 3, "expected 3 Jinx names")

	for i, name := range names {
		require.Equal(t, "Jinx", name, "all names should be Jinx at index %d", i)
	}

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_EmptyResult(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Empty(t, names, "expected empty slice")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_QueryErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(sql.ErrConnDone)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "db query", "error should contain 'db query'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_ScanErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Garen").
		AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "rows scanning", "error should contain 'rows scanning'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetUniqueNames_RowsErrorWrapped(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"})
	rows.AddRow("Garen")
	rows.RowError(0, sql.ErrTxDone)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	service := Mdb.New(db)
	names, err := service.GetUniqueNames()
	require.Error(t, err, "expected error")
	require.Nil(t, names, "expected nil result on error")
	require.Contains(t, err.Error(), "rows error", "error should contain 'rows error'")

	require.NoError(t, mock.ExpectationsWereMet(), "unfulfilled expectations")
}

func TestGetNames_QueryError_NoPanic(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(sql.ErrConnDone)

	service := Mdb.New(db)
	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "db query")

	require.NoError(t, mock.ExpectationsWereMet())
}
