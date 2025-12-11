package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestDBService_GetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("User1"))

	s := New(db)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"User1"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_Empty(t *testing.T) {
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

func TestDBService_GetNames_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrNoRows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("User2"))

	s := New(db)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"User2"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_Empty(t *testing.T) {
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

func TestDBService_GetUniqueNames_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(456)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
