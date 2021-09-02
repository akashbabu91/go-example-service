package main

import (
	"context"
	"example-service/internal/dataaccess"
	"example-service/internal/server"
	"example-service/internal/routes"
	"example-service/internal/logger"
	"fmt"
	"log"
	"net/http"
	"os"
    "os/signal"
	"time"
)

func main() {
	var wait time.Duration
	server, err := server.New()

	if err != nil {
		fmt.Println("Failed to start server ", err)
	}

	employeeReader, err := dataaccess.NewDummy()

	if err != nil {
		fmt.Println("Failed to get dummy", err)
	}

	// register routes
	routes.RegisterRoutes(*server, employeeReader)

	// register middleware
	server.Router.Use(logger.LoggingMiddleware)
	
	srv := &http.Server{
		Handler:      server.Router,
		Addr:         "127.0.0.1:9000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("shutting down")
    os.Exit(0)
}
