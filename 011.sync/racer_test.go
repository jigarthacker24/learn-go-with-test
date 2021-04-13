package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fast url test", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("Error received. error: %v", err)
		}

		if want != got {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("timeout test", func(t *testing.T) {
		server := makeDelayedServer(50 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 25*time.Millisecond)

		if err == nil {
			t.Errorf("Want error but not received")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
