package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	d "github.com/Ekaterina-101/task-6/internal/db"
)

const (
	queryName       = "SELECT name FROM users"
	queryUniqueName = "SELECT DISTINCT name FROM users"
)

var (
	errQueryError = errors.New("query error")
	errDropped    = errors.New("dropped")
)

func TestGetNamesSuccess(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Kamil")

	mock.ExpectQuery(queryName).WillReturnRows(rows)

	want := []string{"Kamil"}
	database := d.New(db)
	have, err := database.GetNames()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesQueryError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryName).WillReturnError(errQueryError)

	database := d.New(db)
	have, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "db query: query error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesScanError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Kamil").
		AddRow(nil)

	mock.ExpectQuery(queryName).WillReturnRows(rows)

	database := d.New(db)
	have, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "rows scanning")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesRowCloseError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errDropped)

	mock.ExpectQuery(queryName).WillReturnRows(rows)

	database := d.New(db)
	have, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "rows error: ")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesSuccess(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Kamil")

	mock.ExpectQuery(queryUniqueName).WillReturnRows(rows)

	want := []string{"Kamil"}
	database := d.New(db)
	have, err := database.GetUniqueNames()

	require.NoError(t, err)
	assert.Equal(t, want, have)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(queryUniqueName).WillReturnError(errQueryError)

	database := d.New(db)
	have, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "db query: query error")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Kamil").
		AddRow(nil)

	mock.ExpectQuery(queryUniqueName).WillReturnRows(rows)

	database := d.New(db)
	have, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "rows scanning")

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesRowCloseError(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errDropped)

	mock.ExpectQuery(queryUniqueName).WillReturnRows(rows)

	database := d.New(db)
	have, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, have)
	assert.Contains(t, err.Error(), "rows error: ")

	require.NoError(t, mock.ExpectationsWereMet())
}
