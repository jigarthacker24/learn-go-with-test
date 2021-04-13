package main

func Sum(in []int) int {
	sum := 0
	for _, num := range in {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {

		tail := []int{}
		if len(slice) > 0 {
			tail = slice[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}
