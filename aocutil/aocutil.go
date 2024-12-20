package aocutil

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseFile[T any](path string, parser func(string) (T, error)) ([]T, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	output := make([]T, 0)
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line, err := parser(scanner.Text())
		if err != nil {
			return nil, err
		}
		output = append(output, line)
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return output, nil
}

func IntLine(line string, sep string) ([]int, error) {
	output := make([]int, 0)
	for _, part := range strings.Split(line, sep) {
		if part == "" {
			continue
		}

		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		output = append(output, n)
	}

	return output, nil
}

func SumSlice(values []int) int {
	output := 0

	for _, v := range values {
		output += v
	}

	return output
}
