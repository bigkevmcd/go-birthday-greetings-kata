package birthday

type EmployeeRepository interface {
	List() ([]Employee, error)
}
