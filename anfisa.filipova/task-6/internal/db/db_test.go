package db_test

import (
	"errors"
	"testing"

	"github.com/Anfisa111/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

const (
	queryDefault = "SELECT name FROM users"
	queryUnique  = "SELECT DISTINCT name FROM users"
)

var (
	errDataBase = errors.New("database error")
	errRow      = errors.New("row error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	t.Run("success with multiple rows", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			AddRow("Bob")
		mock.ExpectQuery(queryDefault).WillReturnRows(rows)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, []string{"Alice", "Bob"}, names)
	})

	t.Run("success with empty result", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"})
		mock.ExpectQuery(queryDefault).WillReturnRows(rows)

		names, err := service.GetNames()

		require.NoError(t, err)
		require.ElementsMatch(t, []string{}, names)
	})

	t.Run("query error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)

		mock.ExpectQuery(queryDefault).
			WillReturnError(errDataBase)

		names, err := service.GetNames()

		require.ErrorContains(t, err, "db query:")
		require.Nil(t, names)
	})

	t.Run("scan error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery(queryDefault).WillReturnRows(rows)

		names, err := service.GetNames()

		require.ErrorContains(t, err, "rows scanning:")
		require.Nil(t, names)
	})

	t.Run("rows error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errRow)

		mock.ExpectQuery(queryDefault).WillReturnRows(rows)

		names, err := service.GetNames()

		require.ErrorContains(t, err, "rows error:")
		require.Nil(t, names)
	})
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	t.Run("distinct with duplicates", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			AddRow("Alice").
			AddRow("Bob")

		mock.ExpectQuery(queryUnique).WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.NoError(t, err)
		require.Equal(t, []string{"Alice", "Alice", "Bob"}, names)
	})

	t.Run("empty result", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"})

		mock.ExpectQuery(queryUnique).WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.NoError(t, err)
		require.ElementsMatch(t, []string{}, names)
	})

	t.Run("query error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)

		mock.ExpectQuery(queryUnique).
			WillReturnError(errDataBase)

		names, err := service.GetUniqueNames()

		require.ErrorContains(t, err, "db query:")
		require.Nil(t, names)
	})

	t.Run("scan error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

		mock.ExpectQuery(queryUnique).WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.ErrorContains(t, err, "rows scanning:")
		require.Nil(t, names)
	})

	t.Run("rows error", func(t *testing.T) {
		t.Parallel()

		mockDB, mock, err := sqlmock.New()

		require.NoError(t, err)

		defer mockDB.Close()

		service := db.New(mockDB)
		rows := sqlmock.NewRows([]string{"name"}).
			AddRow("Alice").
			RowError(0, errRow)

		mock.ExpectQuery(queryUnique).WillReturnRows(rows)

		names, err := service.GetUniqueNames()

		require.ErrorContains(t, err, "rows error:")
		require.Nil(t, names)
	})
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()

	require.NoError(t, err)

	defer mockDB.Close()

	service := db.New(mockDB)

	require.NotNil(t, service)
	require.NotNil(t, service.DB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Test")

	mock.ExpectQuery(queryDefault).WillReturnRows(rows)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Test"}, names)
}
