package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	dbsvc "github.com/Ekaterina-101/task-6/internal/db"
)

func TestGetNames_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

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
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errors.New("query error"))

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query: query error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_EmptyResult(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan").AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErrAfterIteration(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228").
		CloseError(errors.New("row iteration error"))

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Success(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

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
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errors.New("query error"))

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "db query: query error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_EmptyResult(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan").AddRow(nil)
	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErrAfterIteration(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := dbsvc.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228").
		CloseError(errors.New("row iteration error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}
