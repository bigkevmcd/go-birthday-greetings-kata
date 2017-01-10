package birthday

import (
	"reflect"
	"testing"
	"time"
)

func TestFileRepository(t *testing.T) {
	repo := NewFileRepository("employee_data.txt")

	employees, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}

	expected := []Employee{
		{time.Date(1982, 10, 8, 0, 0, 0, 0, time.UTC), "john.doe@foobar.com", "Doe", "John"},
		{time.Date(1975, 3, 11, 0, 0, 0, 0, time.UTC), "mary.ann@foobar.com", "Ann", "Mary"},
	}
	if !reflect.DeepEqual(employees, expected) {
		t.Errorf("unexpected list: got %s, expected %s", employees, expected)
	}
}
