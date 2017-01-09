package birthday

import (
	"testing"
	"time"
)

func TestSendGreetings(t *testing.T) {
	now := time.Date(2017, time.March, 11, 0, 0, 0, 0, time.UTC)
	err := SendGreetings("employee_data.txt", now, "localhost", 25)
	if err != nil {
		t.Error(err)
	}
}
