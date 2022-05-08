package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting")

	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/_ready").HandlerFunc(readyHandler)
	router.Methods(http.MethodGet).Path("/_live").HandlerFunc(liveHandler)
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Path("/version").Handler(versionHandler)
	apiRouter.Use(loggingMiddleware, handlers.CompressHandler, handlers.RecoveryHandler())

	srv := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	var waiter = make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-waiter

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exited")
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func liveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
