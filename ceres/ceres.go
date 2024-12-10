package ceres

type pointDiff struct {
	ydiff, xdiff int
}

func hasXmas(input [][]byte, y, x, ydiff, xdiff int) bool {
	letters := []byte{'M', 'A', 'S'}

	y += ydiff
	x += xdiff
	for _, char := range letters {
		if (x < 0 || x >= len(input[0]) || y < 0 || y >= len(input)) {
			return false
		}

		if (input[y][x] != char) {
			return false
		}

		y += ydiff
		x += xdiff
	}

	return true
}

func countXmasWord(input [][]byte, y, x int) int {
	directions := []pointDiff{
		{-1, -1},
		{-1,  0},
		{-1,  1},
		{ 0, -1},
		{ 0,  1},
		{ 1, -1},
		{ 1,  0},
		{ 1,  1},
	}

	output := 0
	for _, dir := range directions {
		if hasXmas(input, y, x, dir.ydiff, dir.xdiff) {
			output++
		}
	}

	return output
}

func CountXmas(input [][]byte) int {
	output := 0

	for y, row := range input {
		for x, char := range row {
			if char == 'X' {
				output += countXmasWord(input, y, x)
			}
		}
	}

	return output
}

func CountMasX(input [][]byte) int {
	output := 0

	for y, row := range input {
		if y == 0  || y == len(input) - 1 {
			continue
		}
		for x, char := range row {
			if x == 0 || x == len(row) - 1 {
				continue
			}

			if char == 'A'  && isMasX(input, y, x) {
				output++
			}
		}
	}

	return output
}

func sliceDelete[T comparable](slice []T, elem T) []T {
	for i, v := range(slice) {
		if v == elem {
			slice[i] = slice[len(slice) - 1]
			return slice[:len(slice) -1]
		}
	}

	return slice
}

func isMasX(input [][]byte, y, x int) bool {
	corners := []pointDiff{
		{-1, -1},
		{-1,  1},
		{ 1, -1},
		{ 1,  1},
	}

	neighbours := map[pointDiff][]pointDiff {
		{-1, -1}: {{ 1, -1}, {-1,  1}},
		{-1,  1}: {{-1, -1}, { 1,  1}},
		{ 1, -1}: {{-1, -1}, { 1,  1}},
		{ 1,  1}: {{ 1, -1}, {-1,  1}},
	}

	for _, corner := range corners {
		if input[y - corner.ydiff][x - corner.xdiff] == 'M' {
			adj := neighbours[corner]
			for _, n := range adj {
				if input[y - n.ydiff][x - n.xdiff] == 'M' {
					corners = sliceDelete(corners, corner)
					corners = sliceDelete(corners, n)
					for _, c := range corners {
						if input[y - c.ydiff][x - c.xdiff] != 'S' {
							return false
						}
					}

					return true
				}
			}

			return false
		}
	}

	return false
}
