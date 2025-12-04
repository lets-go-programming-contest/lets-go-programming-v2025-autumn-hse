package db_test

import (
	"errors"
	"testing"

	database "github.com/6ermvH/german.feskov/task-6/internal/db"
	"github.com/stretchr/testify/require"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	queryNames       = "SELECT name FROM users"
	queryUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestGetNamesNoErrorOneRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	want := []string{"german"}

	have, err := dbService.GetNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetNamesNoErrorMoreRows(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	want := []string{"german", "anthon", "vitaly"}

	have, err := dbService.GetNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetNamesErrorOnQuery(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	dbService := database.New(db)

	mock.ExpectQuery(queryNames).WillReturnError(errors.ErrUnsupported)

	_, err = dbService.GetNames()
	require.Error(t, err)

	prefix := "db query: "
	require.ErrorContains(t, err, prefix)
}

func TestGetNamesErrorOnScanRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	_, err = dbService.GetNames()
	require.Error(t, err)

	prefix := "rows scanning: "
	require.ErrorContains(t, err, prefix)
}

func TestGetNamesErrorOnCloseRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	_, err = dbService.GetNames()
	require.Error(t, err)

	prefix := "rows error: "
	require.ErrorContains(t, err, prefix)
}

func TestGetUniqueNamesNoErrorOneRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	want := []string{"german"}

	have, err := dbService.GetUniqueNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetUniqueNamesNoErrorMoreRows(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	want := []string{"german", "anthon", "vitaly"}

	have, err := dbService.GetUniqueNames()
	require.NoError(t, err)

	require.Equal(t, want, have)
}

func TestGetUniqueNamesErrorOnQuery(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	mock.ExpectQuery(queryUniqueNames).WillReturnError(errors.ErrUnsupported)

	_, err = dbService.GetUniqueNames()
	require.Error(t, err)

	prefix := "db query: "
	require.ErrorContains(t, err, prefix)
}

func TestGetUniqueNamesErrorOnScanRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	_, err = dbService.GetUniqueNames()
	require.Error(t, err)

	prefix := "rows scanning: "
	require.ErrorContains(t, err, prefix)
}

func TestGetUniqueNamesErrorOnCloseRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	_, err = dbService.GetUniqueNames()
	require.Error(t, err)

	prefix := "rows error: "
	require.ErrorContains(t, err, prefix)
}
