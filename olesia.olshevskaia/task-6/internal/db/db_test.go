package db_test

import (
	"errors"
	"strings"
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

	//names, err := service.GetNames()

	//require.Nil(t, names)

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

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows error")
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

	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesRowsErr(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Hedgehoginthefog").
		RowError(0, errRow)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "rows error") {
		t.Fatalf("unexpected error: %v", err)
	}
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

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "db query")
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

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows scanning")
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

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}
