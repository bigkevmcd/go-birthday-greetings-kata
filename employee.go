package birthday

import (
	"time"
)

type Employee struct {
	Birthdate time.Time
	Email     string
	Firstname string
	Lastname  string
}

func (e Employee) IsBirthday(date time.Time) bool {
	return (e.Birthdate.Month() == date.Month()) && (e.Birthdate.Day() == date.Day())
}

func NewEmployee(firstName, lastName, birthDate, email string) (*Employee, error) {
	date, err := time.Parse("2006/01/02", birthDate)

	if err != nil {
		return nil, err
	}

	return &Employee{
		Firstname: firstName,
		Lastname:  lastName,
		Birthdate: date,
		Email:     email,
	}, nil
}
