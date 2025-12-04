package db_test

import (
	"errors"
	"slices"
	"strings"
	"testing"

	database "github.com/6ermvH/german.feskov/task-6/internal/db"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	queryNames       = "SELECT name FROM users"
	queryUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestGetNamesNoErrorOneRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	want := []string{"german"}

	have, err := dbService.GetNames()
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(have, want) {
		t.Fatalf("slices not equal: have: %v, want: %v", have, want)
	}
}

func TestGetNamesNoErrorMoreRows(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	want := []string{"german", "anthon", "vitaly"}

	have, err := dbService.GetNames()
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(have, want) {
		t.Fatalf("slices not equal: have: %v, want: %v", have, want)
	}
}

func TestGetNamesErrorOnQuery(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	mock.ExpectQuery(queryNames).WillReturnError(errors.ErrUnsupported)

	_, err = dbService.GetNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "db query: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}

func TestGetNamesErrorOnScanRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	_, err = dbService.GetNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "rows scanning: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}

func TestGetNamesErrorOnCloseRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
	mock.ExpectQuery(queryNames).WillReturnRows(rows)

	_, err = dbService.GetNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "rows error: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}

func TestGetUniqueNamesNoErrorOneRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	want := []string{"german"}

	have, err := dbService.GetUniqueNames()
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(have, want) {
		t.Fatalf("slices not equal: have: %v, want: %v", have, want)
	}
}

func TestGetUniqueNamesNoErrorMoreRows(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	want := []string{"german", "anthon", "vitaly"}

	have, err := dbService.GetUniqueNames()
	if err != nil {
		t.Fatalf("return error: %s", err.Error())
	}

	if !slices.Equal(have, want) {
		t.Fatalf("slices not equal: have: %v, want: %v", have, want)
	}
}

func TestGetUniqueNamesErrorOnQuery(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	mock.ExpectQuery(queryUniqueNames).WillReturnError(errors.ErrUnsupported)

	_, err = dbService.GetUniqueNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "db query: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}

func TestGetUniqueNamesErrorOnScanRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	_, err = dbService.GetUniqueNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "rows scanning: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}

func TestGetUniqueNamesErrorOnCloseRow(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := database.New(db)

	rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
	mock.ExpectQuery(queryUniqueNames).WillReturnRows(rows)

	_, err = dbService.GetUniqueNames()
	if err == nil {
		t.Fatalf("return no error")
	}

	prefix := "rows error: "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
	}
}
