package historian

func DiffLists(l0, l1 []int) []int {
	output := make([]int, len(l0))

	for i, v0 := range(l0) {
		output[i] = intAbs(l1[i] - v0)
	}

	return output
}

func SimilarityScore(left, right []int) int {
	counts := make(map[int]int)
	score := 0

	for _, v := range right {
		counts[v] += v
	}

	for _, v := range left {
		score += counts[v]
	}

	return score
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func SumSlice(s []int) int {
	output := 0
	for _, v := range(s) {
		output += v
	}

	return output
}
