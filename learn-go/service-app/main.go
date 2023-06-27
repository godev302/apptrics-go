package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-app/handlers"
	"time"
)

func main() {
	log := log.New(os.Stdout, "users: ", log.LstdFlags|log.Lshortfile)
	err := startApp(log)
	if err != nil {
		log.Fatalln(err)
	}
}

// startApp set up dependencies for the project
func startApp(log *log.Logger) error {

	log.Println("main : Started : initializing handlers support")
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		Handler:      handlers.Api(), // register handler functions
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt) // this will notify the shutdown chan if someone presses ctr+c

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		//and then waiting indefinitely for connections to return to idle and then shut down.
		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil
}
