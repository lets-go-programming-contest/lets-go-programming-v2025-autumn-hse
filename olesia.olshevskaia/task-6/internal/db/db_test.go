package db_test

import (
	"errors"
	"testing"

	"github.com/OlesiaOl/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNamesOK(t *testing.T) {
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
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errors.New("connection lost"))

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "db query")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesRowsError(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	service := db.New(sqlDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, errors.New("row broken"))

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}
