package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/VlasfimosY/task-6/internal/db"
)

var errClose = errors.New("close error")
var errConnDone = errors.New("connection done")

func TestGetNames(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice"))

	s := db.New(dbConn)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := db.New(dbConn)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrNoRows)

	s := db.New(dbConn)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	s := db.New(dbConn)
	_, err = s.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Bob"))

	s := db.New(dbConn)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Bob"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Empty(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := db.New(dbConn)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errConnDone)

	s := db.New(dbConn)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	s := db.New(dbConn)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErr(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").CloseError(errClose)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	s := db.New(dbConn)
	_, err = s.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErr(t *testing.T) {
	t.Parallel()
	dbConn, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Bob").CloseError(errClose)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	s := db.New(dbConn)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}
