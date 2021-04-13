package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleepWord = "sleep"
const writeWord = "write"

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleepWord)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, writeWord)
	return
}

type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buff := bytes.Buffer{}

		cndnOpSpy := &CountdownOperationSpy{}

		Countdown(&buff, cndnOpSpy)
		got := buff.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

	})
	t.Run("sleep before every print", func(t *testing.T) {

		cndnOpSpy := &CountdownOperationSpy{}

		Countdown(cndnOpSpy, cndnOpSpy)
		want := []string{
			sleepWord,
			writeWord,
			sleepWord,
			writeWord,
			sleepWord,
			writeWord,
			sleepWord,
			writeWord,
		}
		if !reflect.DeepEqual(want, cndnOpSpy.Calls) {
			t.Errorf("got: %q, want: %q", cndnOpSpy.Calls, want)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime1 := &spyTime{}

	sleeper := &ConfigurableSleeper{duration: sleepTime, sleep: spyTime1.Sleep}

	sleeper.Sleep()

	if spyTime1.durationSlept != sleepTime {
		t.Errorf("Sleep time not matched. want: %v, got %v", sleepTime, spyTime1.durationSlept)
	}
}

func TestMyMain(t *testing.T) {
	MyMain()
}
