package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/OlesiaOl/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var (
	errLostConnection = errors.New("connection lost")
	errRowBroken      = errors.New("row broken")
	errRow            = errors.New("some rows error")
)

func TestGetNamesOK(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("OlesyainWonderland").
		AddRow("Olesyainthelandofnightmares")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"OlesyainWonderland", "Olesyainthelandofnightmares"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesQueryError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errLostConnection)

	_, err = service.GetNames()

	require.ErrorContains(t, err, "db query")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesRowsError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, errRowBroken)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetNames()

	require.ErrorContains(t, err, "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesScanError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetNames()

	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesRowsErr(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Hedgehoginthefog").
		RowError(0, errRow)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetUniqueNames()

	require.ErrorContains(t, err, "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesOK(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Tom").
		AddRow("Jerry")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Tom", "Jerry"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errLostConnection)

	_, err = service.GetUniqueNames()

	require.ErrorContains(t, err, "db query")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetUniqueNames()

	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesRowsError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Aliceandmyelophone").
		RowError(0, errRow)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetUniqueNames()

	require.ErrorContains(t, err, "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}
