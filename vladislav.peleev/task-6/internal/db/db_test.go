package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNames(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("A"))

	s := New(db)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"A"}, names)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetNamesEmpty(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := New(db)
	names, err := s.GetNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetNamesScanErr(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	r := sqlmock.NewRows([]string{"name"}).AddRow(123)
	m.ExpectQuery("SELECT name FROM users").WillReturnRows(r)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetNamesQueryErr(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT name FROM users").WillReturnError(sql.ErrNoRows)

	s := New(db)
	_, err = s.GetNames()
	require.Error(t, err)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("B"))

	s := New(db)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"B"}, names)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetUniqueNamesEmpty(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	s := New(db)
	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetUniqueNamesScanErr(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	r := sqlmock.NewRows([]string{"name"}).AddRow(456)
	m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(r)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, m.ExpectationsWereMet())
}

func TestGetUniqueNamesQueryErr(t *testing.T) {
	db, m, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(sql.ErrConnDone)

	s := New(db)
	_, err = s.GetUniqueNames()
	require.Error(t, err)
	require.NoError(t, m.ExpectationsWereMet())
}
