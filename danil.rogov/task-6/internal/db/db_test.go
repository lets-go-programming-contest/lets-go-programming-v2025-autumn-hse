package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tapochek2894/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	namesQuery        = "SELECT name FROM users"
	testName          = "Alice"
	rowsError         = "rows error: "
	rowsScanningError = "rows scanning: "
	uniqueNamesQuery  = "SELECT DISTINCT name FROM users"
)

var (
	errClosing = errors.New("error during close")
	errQuery   = errors.New("error during query")
)

//nolint:ireturn
func createTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()

	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	return db, mock
}

func TestCorrectGetNames(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(testName)

	mock.ExpectQuery(namesQuery).WillReturnRows(rows)

	expected := []string{testName}
	database := db.New(mockDatabase)
	got, err := database.GetNames()

	require.NoError(t, err)
	assert.Equal(t, expected, got)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestIncorrectGetNames(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	mock.ExpectQuery(namesQuery).WillReturnError(errQuery)

	database := db.New(mockDatabase)
	got, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, got)
	require.ErrorContains(t, err, rowsError)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesScanError(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(namesQuery).WillReturnRows(rows)

	database := db.New(mockDatabase)
	got, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, got)
	require.ErrorContains(t, err, rowsScanningError)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNamesRowCloseError(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errClosing)

	mock.ExpectQuery(namesQuery).WillReturnRows(rows)

	database := db.New(mockDatabase)
	expected, err := database.GetNames()

	require.Error(t, err)
	assert.Nil(t, expected)
	require.ErrorContains(t, err, rowsError)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCorrectGetUniqueNames(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(testName)

	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(rows)

	expected := []string{testName}
	database := db.New(mockDatabase)
	got, err := database.GetUniqueNames()

	require.NoError(t, err)
	assert.Equal(t, expected, got)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestIncorrectGetUniqueNames(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	mock.ExpectQuery(namesQuery).WillReturnError(errQuery)

	database := db.New(mockDatabase)
	got, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, got)
	require.ErrorContains(t, err, rowsError)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(rows)

	database := db.New(mockDatabase)
	got, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, got)
	require.ErrorContains(t, err, rowsScanningError)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNamesRowCloseError(t *testing.T) {
	t.Parallel()

	mockDatabase, mock := createTestDB(t)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errClosing)

	mock.ExpectQuery(uniqueNamesQuery).WillReturnRows(rows)

	database := db.New(mockDatabase)
	expected, err := database.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, expected)
	require.ErrorContains(t, err, rowsError)

	require.NoError(t, mock.ExpectationsWereMet())
}
