package levenshtein_test

import (
	"testing"

	"github.com/emmanuelay/levenshtein"
)

func TestCalculate(t *testing.T) {

	tt := []struct {
		a, b     string
		distance int
	}{
		{a: "cat", b: "cap", distance: 1},
		{a: "", b: "", distance: 0},
		{a: "1", b: "1", distance: 0},
		{a: "1", b: "2", distance: 1},
		{a: "12", b: "12", distance: 0},
		{a: "123", b: "12", distance: 1},
		{a: "1234", b: "1", distance: 3},
		{a: "1234", b: "1233", distance: 1},
		{a: "1248", b: "1349", distance: 2},
		{a: "", b: "12345", distance: 5},
		{a: "5677", b: "1234", distance: 4},
		{a: "123456", b: "12345", distance: 1},
		{a: "13579", b: "12345", distance: 4},
		{a: "123", b: "", distance: 3},
		{a: "javascript", b: "javscript", distance: 1},
		{a: "javascript", b: "jvscript", distance: 2},
		{a: "javascript", b: "jvscropt", distance: 3},
		{a: "javascript", b: "jvscrop", distance: 4},
		{a: "javascript", b: "jvscri", distance: 4},
		{a: "javascript", b: "jvscro", distance: 5},
		{a: "saluhallen", b: "sluhllen", distance: 2},
		{a: "avenyn", b: "avnyn", distance: 1},
		{a: "avenyn", b: "avny", distance: 2},
		{a: "avenyn", b: "avn", distance: 3},
		{a: "avenyn", b: "av", distance: 4},
		{a: "Avenyn", b: "av", distance: 5},
	}

	for _, test := range tt {
		if val := levenshtein.Distance(test.a, test.b); val != test.distance {
			t.Errorf("Expected %d, got %d", test.distance, val)
		}
	}
}
