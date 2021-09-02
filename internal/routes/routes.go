package routes

import (
	"example-service/internal/server"
	"example-service/internal/dataaccess"
	"example-service/internal/service"
	"fmt"
	"encoding/json"
	"net/http"
)

func RegisterRoutes(server server.Server, readerWriter dataaccess.ReaderWriter) {
	server.Router.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling api/test")
		employees, err := service.GetAllEmployees(readerWriter)

		if err != nil {
			fmt.Println("Failed to read employee", err)
		}

		// an example API handler
		json.NewEncoder(w).Encode(employees)
	})
}