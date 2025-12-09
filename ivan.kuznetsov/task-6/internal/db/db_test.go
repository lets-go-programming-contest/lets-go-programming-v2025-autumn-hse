package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kuzid-17/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var (
	errExpected = errors.New("ExpectedError")
	errRows     = errors.New("rows error")
)

type rowTestDB struct {
	names       []string
	errExpected error
	scanError   bool
	rowsError   bool
}

var testTable = []rowTestDB{
	{
		names: []string{"Ivan", "Gena228"},
	},
	{
		names:       nil,
		errExpected: errExpected,
	},
	{
		names:     []string{"Ivan"},
		scanError: true,
	},
	{
		names:     []string{"Ivan", "Gena228"},
		rowsError: true,
	},
}

func TestNew(t *testing.T) {
	t.Parallel()
	mockDB, _, _ := sqlmock.New()
	dbService := db.New(mockDB)

	require.NotNil(t, dbService, "service should not be nil")
	require.Equal(t, mockDB, dbService.DB, "DB field should be set correctly")
}

func TestGetName(t *testing.T) {
	t.Parallel()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range testTable {
		rows := mockDBRows(row)
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows).WillReturnError(row.errExpected)

		names, err := dbService.GetNames()

		if (row.errExpected != nil) || row.scanError || row.rowsError {
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestGetUniqueName(t *testing.T) {
	t.Parallel()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.DBService{DB: mockDB}

	for i, row := range testTable {
		rows := mockDBRows(row)
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows).WillReturnError(row.errExpected)

		values, err := dbService.GetUniqueNames()

		if (row.errExpected != nil) || row.scanError || row.rowsError {
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, values, "row: %d, expected values: %s, actual values: %s", i, row.names, values)
	}
}

func mockDBRows(row rowTestDB) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})

	if row.scanError {
		rows = rows.AddRow(nil)
	}

	for _, name := range row.names {
		rows = rows.AddRow(name)
	}

	if row.rowsError {
		rows = rows.RowError(0, errRows)
	}

	return rows
}
