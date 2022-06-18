package main

import (
	"log"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
)

func loggingMiddleware(h http.Handler) http.Handler {
	return loggingHandler{handler: h}
}

type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t := time.Now()

	status := http.StatusOK

	wrapped := httpsnoop.Wrap(w, httpsnoop.Hooks{
		WriteHeader: func(httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
			return func(code int) {
				status = code
			}
		},
	})

	h.handler.ServeHTTP(wrapped, req)

	log.Println(req.Method, req.URL, time.Now().Sub(t), status, req.Header)
}
