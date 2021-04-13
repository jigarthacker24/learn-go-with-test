package iteration

import "strings"

// Repeat takes input string and count and generates new string with repeating the count times.
func Repeat(in string, count int) string {
	out := ""

	if count < 0 {
		return out
	}
	return strings.Repeat(in, count)

}
