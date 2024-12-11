package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"atmosia.net/aoc2024/aocutil"
	"atmosia.net/aoc2024/printqueue"
)

func lineParser(line string) ([]int, error) {
	if line == "" {
		return nil, nil
	}

	if strings.Contains(line, "|") {
		return aocutil.IntLine(line, "|")
	}

	return aocutil.IntLine(line, ",")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	lines, err := aocutil.ParseFile(os.Args[1], lineParser)
	if err != nil {
		panic(err)
	}

	idx := slices.IndexFunc(lines, func(e []int) bool { return e == nil })
	rules := buildRules(lines[:idx])
	updates := lines[idx+1:]

	output := 0
	invalidOrder := make([][]int, 0)
	for _, update := range updates {
		if printqueue.ValidOrdering(update, rules) {
			output += printqueue.FindMiddle(update)
		} else {
			invalidOrder = append(invalidOrder, update)
		}
	}

	fmt.Printf("%v\n", output)

	output = 0
	for _, update := range invalidOrder {
		reordered := printqueue.Reorder(update, rules)
		output += printqueue.FindMiddle(reordered)
	}

	fmt.Printf("%v\n", output)
}

func buildRules(lines [][]int) map[int][]int {
	pairs := make([]printqueue.Pair, 0, len(lines))

	for _, values := range lines {
		pairs = append(pairs, printqueue.Pair{values[0], values[1]})
	}

	return printqueue.BuildOrderingRules(pairs)
}
