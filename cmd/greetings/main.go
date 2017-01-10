package main

import (
	"time"

	"github.com/bigkevmcd/go-birthday-greetings-kata"
)

func main() {
	birthday.SendGreetings("employee_data.txt", time.Now(), "localhost", 25)
}
