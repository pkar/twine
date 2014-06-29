package twine

import (
	"testing"
)

var levMinTests = []struct {
	args []int
	min  int
}{
	{[]int{1, 2, 0}, 0},
	{[]int{}, -1},
	{[]int{5, 1, 9}, 1},
	{[]int{5}, 5},
}

func TestLevMin(t *testing.T) {
	for _, tt := range levMinTests {
		min := levMin(tt.args...)
		if min != tt.min {
			t.Errorf("levMin(%v) => %d, want %d", tt.args, min, tt.min)
		}
	}
}

var levTests = []struct {
	source   string
	target   string
	distance int
}{
	// two empty
	{"", "", 0},
	// deletion
	{"library", "librar", 1},
	// one empty, left
	{"", "library", 7},
	// one empty, right
	{"library", "", 7},
	{"car", "cars", 1},
	{"", "a", 1},
	{"a", "aa", 1},
	{"a", "aaa", 2},
	{"", "", 0},
	{"a", "b", 1},
	{"aaa", "aba", 1},
	{"aaa", "ab", 2},
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"a", "", 1},
	{"aa", "a", 1},
	{"aaa", "a", 2},
	// unicode
	{"Schüßler", "Schübler", 1},
	{"Schüßler", "Schußler", 1},
	{"Schüßler", "Schüßler", 0},
	{"Schüßler", "Schüler", 1},
	{"Schüßler", "Schüßlers", 1},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, tt := range levTests {
		res := LevenshteinDistance(tt.source, tt.target)
		if res != tt.distance {
			t.Errorf("LevenshteinDistance(%s, %s) => %d, want %d", tt.source, tt.target, res, tt.distance)
		}
	}
}
