package reader

func (dummyDataReader DummyDataReader) Read() (*Employee, error) {
	return &Employee{
		Name: "Test",
		Age: 30,
		Company: "Clarabridge",
		Designation: "IT",
	}, nil
}