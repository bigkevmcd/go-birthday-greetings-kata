package main

import (
	"time"

	"github.com/bigkevmcd/birthday-greetings-go"
)

func main() {
	birthday.SendGreetings("employee_data.txt", time.Now(), "localhost", 25)
}
