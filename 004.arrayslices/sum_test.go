package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum of fixed length array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		sum := Sum(nums)
		want := 15
		if sum != want {
			t.Errorf("want: '%d', got: '%d'. Data: %v", want, sum, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of all slices", func(t *testing.T) {
		sums := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{6, 15}
		if !reflect.DeepEqual(sums, want) {
			t.Errorf("want: '%v', got: '%v'", want, sums)
		}
	})

}

func TestSumAllTrails(t *testing.T) {
	checkSum := func(t testing.TB, got, wnt []int) {
		t.Helper()
		if !reflect.DeepEqual(got, wnt) {
			t.Errorf("got: %v, want: %v", got, wnt)
		}
	}

	t.Run("Sum of all slices tails", func(t *testing.T) {
		sums := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}, []int{3, 4, 5})
		want := []int{5, 14, 9}
		checkSum(t, sums, want)
	})

	t.Run("Sum of slices tails with empty slices", func(t *testing.T) {
		sums := SumAllTails([]int{3, 4, 5}, []int{})
		want := []int{9, 0}
		checkSum(t, sums, want)

	})
}

func ExampleSumAllTails() {
	sums := SumAllTails([]int{2, 3, 4}, []int{5, 5, 5}, []int{})
	fmt.Println(sums)
	//Output: [7 10 0]
}
