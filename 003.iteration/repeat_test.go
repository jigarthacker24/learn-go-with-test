package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("repeat 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		wnt := "aaa"

		if got != wnt {
			t.Errorf("got: %q, want: %q", got, wnt)
		}

	})

	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		wnt := "aaaaa"

		if got != wnt {
			t.Errorf("got: %q, want: %q", got, wnt)
		}

	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A", 10)
	}
}

func ExampleRepeat() {
	out := Repeat("J", 4)
	fmt.Println(out)
	//Output: JJJJ
}
