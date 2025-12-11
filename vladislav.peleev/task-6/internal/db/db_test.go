package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice"))

	s := New(db)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := New(db)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrNoRows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Bob"))

	s := New(db)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Bob"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := New(db)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows scanning:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").CloseError(errors.New("close error"))
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Bob").CloseError(errors.New("close error"))
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.Contains(t, err.Error(), "rows error:")
	require.NoError(t, mock.ExpectationsWereMet())
}
