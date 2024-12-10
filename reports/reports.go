package reports

func remove(slice []int, idx int) []int {
	output := make([]int, 0, len(slice)-1)
	output = append(output, slice[:idx]...)
	output = append(output, slice[idx+1:]...)

	return output
}

func isSafe(line []int, dampen bool) bool {
	increments := true
	if line[1] < line[0] {
		increments = false
	}

	if dampen && isSafe(line[1:], false) {
		return true
	}

	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		if increments {
			if diff < 1 || diff > 3 {
				if !dampen {
					return false
				}

				prev := remove(line, i-1)
				next := remove(line, i)
				return isSafe(prev, false) || isSafe(next, false)
			}
		} else {
			if diff > -1 || diff < -3 {
				if !dampen {
					return false
				}

				prev := remove(line, i-1)
				next := remove(line, i)
				return isSafe(prev, false) || isSafe(next, false)
			}
		}
	}

	return true
}

func CountSafe(lines [][]int, dampen bool) int {
	safe := 0

	for _, line := range lines {
		if isSafe(line, dampen) {
			safe += 1
		}
	}

	return safe
}
