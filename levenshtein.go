package twine

// levMin returns the minimum of the given variadic input.
// if no input is given a -1 is returned.
func levMin(args ...int) int {
	length := len(args)
	switch length {
	case 0:
		return -1
	default:
		min := args[0]
		for i := 1; i < length; i++ {
			if args[i] < min {
				min = args[i]
			}
		}
		return min
	}
}

// LevenshteinDistance measures the distance between two strings.
// http://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(source string, target string) int {
	// degenerate cases
	if source == target {
		return 0
	}
	if len(source) == 0 {
		return len(target)
	}
	if len(target) == 0 {
		return len(source)
	}
	t1 := []rune(source)
	t2 := []rune(target)

	// create two work vectors of integer distances
	v0 := make([]int, len(t2)+1)
	v1 := make([]int, len(t2)+1)

	// initialize v0 (the previous row of distances)
	// this row is A[0][i]: edit distance for an empty s
	// the distance is just the number of characters to delete from t
	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < len(t1); i++ {
		// calculate v1 (current row distances) from the previous row v0
		// first element of v1 is A[i+1][0]
		//   edit distance is delete (i+1) chars from s to match empty t
		v1[0] = i + 1

		// use formula to fill in the rest of the row
		var cost int
		for j := 0; j < len(t2); j++ {
			if t1[i] == t2[j] {
				cost = 0
			} else {
				cost = 1
			}
			v1[j+1] = levMin(v1[j]+1, v0[j+1]+1, v0[j]+cost)
		}

		// copy v1 (current row) to v0 (previous row) for next iteration
		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}
	}

	return v1[len(t2)]
}
