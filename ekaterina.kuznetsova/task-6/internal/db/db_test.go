package db_test

import (
	"errors"
	"testing"

	"github.com/Ekaterina-101/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetNames_Success(t *testing.T) {
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

func TestGetNames_EmptyResult(t *testing.T) {
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

func TestGetNames_QueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(errors.New("query error"))

	service := db.New(mockDB)
	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsErrAfterIteration(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan")
	rows.RowError(0, errors.New("iteration error"))

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_CloseErrorOnEmptyRows(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.New("close error"))

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Success(t *testing.T) {
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

func TestGetUniqueNames_EmptyResult(t *testing.T) {
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

func TestGetUniqueNames_QueryError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errors.New("query error"))

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsErrAfterIteration(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Ivan")
	rows.RowError(0, errors.New("iteration error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_CloseErrorOnEmptyRows(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.New("close error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}


func TestNewDBService(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	require.NotNil(t, service)
	require.Equal(t, mockDB, service.DB)
}
