package service

import (
	"example-service/internal/dataaccess"
	"fmt"
)

// function to get all employees using a Reader interface
func GetAllEmployees(reader dataaccess.Reader) ([]dataaccess.Employee, error) {
	employees, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error while retrieving all employees", err)
		return nil, err
	}

	return employees, err
}