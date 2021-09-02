package dataaccess

// Struct to represent an Employee
type Employee struct {
	Name string
	Age int
	Company string
	Designation string
}

type Reader interface {
	Read() (*Employee, error)
	ReadAll() ([]Employee, error)
}

type Writer interface {
	Write() error
}

type ReaderWriter interface {
	Reader
	Writer
}