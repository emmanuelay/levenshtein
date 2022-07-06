package levenshtein

import (
	"math"
)

// Distance returns the Levenshtein distance between two strings
func Distance(a, b string) int {

	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	if len(a) > len(b) {
		tmp := a
		a = b
		b = tmp
	}

	row := make([]int, len(a)+1)

	min := func(val ...int) int {
		m := math.MaxInt
		for _, v := range val {
			if v < m {
				m = v
			}
		}
		return m
	}

	for i := 1; i <= len(b); i++ {
		prev := i
		for j := 1; j <= len(a); j++ {
			var val int
			if b[i-1] == a[j-1] {
				val = row[j-1]
			} else {
				val = min(int(row[j-1]+1),
					int(prev+1),
					int(row[j]+1))
			}
			row[j-1] = prev
			prev = val
		}
		row[len(a)] = prev
	}

	return row[len(a)]
}
