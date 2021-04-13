package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err == nil {
			fmt.Fprintf(rw, data)
		} else {
			//just log and do not write response if error
		}
	}
}
