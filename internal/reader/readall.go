package reader

// slice with dummy values - simulate a DB
var employees = []Employee{
	{Name: "Test-1", Age: 30, Company: "Clarabridge", Designation: "IT"},
	{Name: "Test-2", Age: 25, Company: "Clarabridge", Designation: "HR"},
	{Name: "Test-3", Age: 28, Company: "New Relic", Designation: "IT"},
}

func (dummyDataReader DummyDataReader) ReadAll() ([]Employee, error) {
	return employees, nil
}