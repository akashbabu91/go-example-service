package dataaccess

type dummy struct {
	// dummy Reader returning hardcoded values
}

func NewDummy() (ReaderWriter, error) {
	return config{
		ReaderWriter: dummy{},
	}, nil
}

func NewDummyReader() (Reader, error) {
	return dummy{}, nil
}

func NewDummyWriter() (Reader, error) {
	return dummy{}, nil
}

// slice with dummy values - simulate a DB
var employees = []Employee{
	{Name: "Test-1", Age: 30, Company: "Clarabridge", Designation: "IT"},
	{Name: "Test-2", Age: 25, Company: "Clarabridge", Designation: "HR"},
	{Name: "Test-3", Age: 28, Company: "New Relic", Designation: "IT"},
}

func (d dummy) ReadAll() ([]Employee, error) {
	return employees, nil
}

func (d dummy) Read() (*Employee, error) {
	return &Employee{
		Name: "Test",
		Age: 30,
		Company: "Clarabridge",
		Designation: "IT",
	}, nil
}

func (d dummy) Write() error {
	return nil
}