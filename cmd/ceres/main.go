package main

import (
	"fmt"
	"os"

	"atmosia.net/aoc2024/aocutil"
	"atmosia.net/aoc2024/ceres"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	parser := func(line string) ([]byte, error) {
		return []byte(line), nil
	}
	xmasMap, err := aocutil.ParseFile(os.Args[1], parser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", ceres.CountXmas(xmasMap))
	fmt.Printf("%v\n", ceres.CountMasX(xmasMap))
}
