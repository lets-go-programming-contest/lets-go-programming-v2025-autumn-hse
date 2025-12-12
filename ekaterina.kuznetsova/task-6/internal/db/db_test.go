package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Ekaterina-101/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var (
	errQuery   = errors.New("query error")
	errRowIter = errors.New("row iteration error")
)

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228")

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Gena228"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errQuery)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query: query error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_EmptyResult(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErrAfterIteration(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228").
		RowError(1, errRowIter)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error: row iteration error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Gena228"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errQuery)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query: query error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_EmptyResult(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErrAfterIteration(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228").
		RowError(1, errRowIter)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error: row iteration error")
	require.NoError(t, mock.ExpectationsWereMet())
}
