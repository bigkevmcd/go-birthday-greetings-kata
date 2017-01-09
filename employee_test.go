package birthday

import (
	"testing"
	"time"
)

func TestBirthday(t *testing.T) {
	employee, err := NewEmployee("foo", "bar", "1990/01/31", "a@b.c")
	if err != nil {
		t.Fatal(err)
	}
	date1, err := time.Parse("2006/01/02", "2008/01/30")
	if err != nil {
		t.Fatal(err)
	}
	date2, err := time.Parse("2006/01/02", "2008/01/31")
	if err != nil {
		t.Fatal(err)
	}

	if employee.IsBirthday(date1) {
		t.Errorf("%s is incorrectly identified as the employee birthday", date1)
	}

	if !employee.IsBirthday(date2) {
		t.Errorf("%s was not identified as the employee birthday", date2)
	}
}
