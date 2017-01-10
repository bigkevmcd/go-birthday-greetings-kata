package birthday

import (
	"encoding/csv"
	"io"
	"os"
)

type FileEmployeeRepository struct {
	filename string
}

func (repo FileEmployeeRepository) List() ([]Employee, error) {
	f, err := os.Open(repo.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	employees := make([]Employee, 0)

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	header := false
	for {
		rec, err := r.Read()
		if err == io.EOF {
			return employees, nil
		}
		if err != nil {
			return nil, err
		}
		if !header {
			header = true
			continue
		}
		// Lastname, Firstname, dateOfBirth, email
		e, err := NewEmployee(rec[0], rec[1], rec[2], rec[3])
		if err != nil {
			return nil, err
		}
		employees = append(employees, *e)
	}

	return employees, nil
}

func NewFileRepository(filename string) FileEmployeeRepository {
	return FileEmployeeRepository{filename: filename}
}
