package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	wnt := 4

	if sum != wnt {
		t.Errorf("Sum test. Got: '%d', want: '%d'", sum, wnt)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output:6
}
