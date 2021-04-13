package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("inc counter 3 times", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
	t.Run("safe concurrently", func(t *testing.T) {

		wantCount := 10000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCount(t, counter, wantCount)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got: %d, want: %d", got.Value(), want)
	}
}
