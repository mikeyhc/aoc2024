package main

import (
	"fmt"
	"os"

	"atmosia.net/aoc2024/aocutil"
	"atmosia.net/aoc2024/reports"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	parser := func(line string) ([]int, error) {
		return aocutil.IntLine(line)
	}
	lines, err := aocutil.ParseFile(os.Args[1], parser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", reports.CountSafe(lines, false))
	fmt.Printf("%v\n", reports.CountSafe(lines, true))
}
