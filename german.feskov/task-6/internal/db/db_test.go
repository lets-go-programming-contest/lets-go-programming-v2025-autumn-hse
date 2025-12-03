//nolint: testpackage
package db

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

//nolint: paralleltest
func TestGetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := New(db)

	query := "SELECT name FROM users"

	t.Run("no error, one row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
		mock.ExpectQuery(query).WillReturnRows(rows)

		want := []string{"german"}

		have, err := dbService.GetNames()
		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		if !slices.Equal(have, want) {
			t.Fatalf("slices not equal: have: %v, want: %v", have, want)
		}
	})

	t.Run("no error, more rows", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
		mock.ExpectQuery(query).WillReturnRows(rows)

		want := []string{"german", "anthon", "vitaly"}

		have, err := dbService.GetNames()
		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		if !slices.Equal(have, want) {
			t.Fatalf("slices not equal: have: %v, want: %v", have, want)
		}
	})

	t.Run("error on query", func(t *testing.T) {
		mock.ExpectQuery(query).WillReturnError(errors.ErrUnsupported)

		_, err := dbService.GetNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "db query: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})

	t.Run("error on scan row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := dbService.GetNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "rows scanning: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})

	t.Run("error on close row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := dbService.GetNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "rows error: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})
}

//nolint: paralleltest
func TestGetUniqueNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("initialize mock db: %s", err.Error())
	}

	dbService := New(db)

	query := "SELECT DISTINCT name FROM users"

	t.Run("no error, one row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("german")
		mock.ExpectQuery(query).WillReturnRows(rows)

		want := []string{"german"}

		have, err := dbService.GetUniqueNames()
		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		if !slices.Equal(have, want) {
			t.Fatalf("slices not equal: have: %v, want: %v", have, want)
		}
	})

	t.Run("no error, more rows", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("german").AddRow("anthon").AddRow("vitaly")
		mock.ExpectQuery(query).WillReturnRows(rows)

		want := []string{"german", "anthon", "vitaly"}

		have, err := dbService.GetUniqueNames()
		if err != nil {
			t.Fatalf("return error: %s", err.Error())
		}

		if !slices.Equal(have, want) {
			t.Fatalf("slices not equal: have: %v, want: %v", have, want)
		}
	})

	t.Run("error on query", func(t *testing.T) {
		mock.ExpectQuery(query).WillReturnError(errors.ErrUnsupported)

		_, err := dbService.GetUniqueNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "db query: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})

	t.Run("error on scan row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := dbService.GetUniqueNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "rows scanning: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})

	t.Run("error on close row", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"name"}).CloseError(errors.ErrUnsupported)
		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := dbService.GetUniqueNames()
		if err == nil {
			t.Fatalf("return no error")
		}

		prefix := "rows error: "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Fatalf("error: %q, don't has prefix: %q", err.Error(), prefix)
		}
	})
}
