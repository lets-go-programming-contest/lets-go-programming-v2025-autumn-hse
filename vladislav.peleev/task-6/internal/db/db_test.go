package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDBService_GetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").AddRow("Bob").AddRow("Charlie"))

	service := New(db)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob", "Charlie"}, names)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	service := New(db)

	names, err := service.GetNames()
	require.NoError(t, err)
	assert.Empty(t, names)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(sql.ErrNoRows)

	service := New(db)

	_, err = service.GetNames()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "db query:")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetNames_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name"}).AddRow(123) // wrong type
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	service := New(db)

	_, err = service.GetNames()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rows scanning:")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").AddRow("Bob"))

	service := New(db)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{"Alice", "Bob"}, names)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBService_GetUniqueNames_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(sql.ErrConnDone)

	service := New(db)

	_, err = service.GetUniqueNames()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "db query:")

	assert.NoError(t, mock.ExpectationsWereMet())
}
