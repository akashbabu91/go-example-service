package reader

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

type S3DataReader struct {
	// things needed to make S3DataReader functional
}

type DummyDataReader struct {
	// dummy Reader returning hardcoded values
}

func NewDummyDataReader() (*DummyDataReader, error) {
	return &DummyDataReader{}, nil
}
