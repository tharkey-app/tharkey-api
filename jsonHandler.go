package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func jsonHandler[T any](f func(context.Context) (T, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		res, err := f(ctx)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(&res); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})
}
