package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("Not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	go func() {
		var result string

		for _, chr := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(chr)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func TestServer(t *testing.T) {

	t.Run("returns the data", func(t *testing.T) {
		data := "Hello, Jigar"
		store := &SpyStore{response: data, t: t}

		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("want: %q, got: %q", data, response.Body.String())
		}

	})

	t.Run("tells store to cancel the request", func(t *testing.T) {
		data := "Hello, Jigar"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Errorf("Response should not have been written")
		}
	})
}
