package service

import (
	"example-service/internal/reader"
	"fmt"
)

// function to get all employees using a Reader interface
func GetAllEmployees(reader reader.Reader) ([]reader.Employee, error) {
	employees, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error while retrieving all employees", err)
	}

	return employees, err
}