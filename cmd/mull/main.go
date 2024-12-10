package main

import (
	"fmt"
	"os"

	"atmosia.net/aoc2024/mull"
	"atmosia.net/aoc2024/aocutil"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	parser := func(line string) (string, error) {
		return line, nil
	}
	lines, err := aocutil.ParseFile(os.Args[1], parser)
	if err != nil {
		panic(err)
	}

	commands := make([]mull.Command, 0)
	for _, line := range lines {
		 commands = append(commands, mull.Parse([]byte(line))...)
	}

	values, err := mull.RunCommands(commands, false)
	if err != nil {
		panic(err)
	}
	sum := aocutil.SumSlice(values)
	fmt.Printf("%v\n", sum)

	values, err = mull.RunCommands(commands, true)
	if err != nil {
		panic(err)
	}
	sum = aocutil.SumSlice(values)
	fmt.Printf("%v\n", sum)
}
