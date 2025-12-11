package db_test

import (
	"errors"
	"github.com/Ekaterina-101/task-6/internal/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names: []string{"Ivan", "Gena228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}


func TestGetNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
}

func TestGetNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("row error"))
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := dbService.GetNames()
	require.Error(t, err)
	require.Nil(t, names)
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}
	return rows
}

type uniqueNamesTestCase struct {
	names       []string
	errExpected error
}

var uniqueNamesTestTable = []uniqueNamesTestCase{
	{
		names: []string{"Ivan", "Gena228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetUniqueNames_Success(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Gena228")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := dbService.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Gena228"}, names)
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errors.New("query error"))

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dbService := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("row error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := dbService.GetUniqueNames()
	require.Error(t, err)
	require.Nil(t, names)
}


func mockUniqueRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}
	return rows
}
