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

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing sqlmock", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range testTable {
		mock.
			ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDbRows(row.names)).
			WillReturnError(row.errExpected)

		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %v, actual error: %v", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %v, actual names: %v", i, row.names, names)
	}
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

func TestGetUniqueNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock init error: %s", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range uniqueNamesTestTable {

		if row.errExpected != nil {
			mock.
				ExpectQuery("SELECT DISTINCT name FROM users").
				WillReturnError(row.errExpected)
		} else {
			mock.
				ExpectQuery("SELECT DISTINCT name FROM users").
				WillReturnRows(mockUniqueRows(row.names))
		}

		names, err := dbService.GetUniqueNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d", i)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d", i)
		require.Equal(t, row.names, names, "row: %d", i)
	}

}

func mockUniqueRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}
	return rows
}
