package printqueue

type Pair struct { Left, Right int }

func BuildOrderingRules(input []Pair) map[int][]int {
	output := make(map[int][]int)

	for _, pair := range input {
		output[pair.Left] = append(output[pair.Left], pair.Right)
	}

	return output
}

func ValidOrdering(input []int, rules map[int][]int) bool {
	seen := make(map[int]bool)

	for _, n := range input {
		pages, ok := rules[n]
		if ok {
			for _, p := range pages {
				_, ok := seen[p]
				if ok {
					return false
				}
			}
		}

		seen[n] = true
	}

	return true
}

func FindMiddle(input []int) int {
	return input[len(input) / 2]
}

func Reorder(input []int, rules map[int][]int) []int {
	seen := make(map[int]int)

	output := make([]int, len(input))
	for i, update := range input {
		if idx := seenAt(rules[update], seen); idx >= 0 {
			shuffle(output, seen, idx, i)
			output[idx] = update
			seen[update] = idx
		} else {
			seen[update] = i
			output[i] = update
		}
	}

	return output
}

func seenAt(rules []int, seen map[int]int) int {
	idx := -1

	for _, rule := range rules {
		if i, ok := seen[rule]; ok {
			if idx < 0 || i < idx {
				idx = i
			}
		}
	}

	return idx
}

func shuffle(output []int, seen map[int]int, idx int, end int) {
	for i := end; i > idx; i-- {
		seen[output[i]] = seen[output[i]] + 1
		output[i] = output[i - 1]
	}
	seen[output[idx]] = seen[output[idx]] + 1
}
