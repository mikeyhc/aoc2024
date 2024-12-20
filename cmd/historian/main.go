package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"atmosia.net/aoc2024/historian"
)

func readList(file string) ([]int, []int) {
	l0, l1 := make([]int, 0), make([]int, 0)
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		left, right, _ := strings.Cut(line, " ")
		right = strings.TrimSpace(right)
		n0, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}
		n1, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}
		l0 = append(l0, n0)
		l1 = append(l1, n1)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	sortFunc := func(slice []int) func(int, int) bool {
		return func(i, j int) bool {
			return slice[i] < slice[j]
		}
	}

	sort.Slice(l0, sortFunc(l0))
	sort.Slice(l1, sortFunc(l1))

	return l0, l1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	list1, list2 := readList(os.Args[1])
	diffList := historian.DiffLists(list1, list2)
	fmt.Printf("%v\n", historian.SumSlice(diffList))
	fmt.Printf("%v\n", historian.SimilarityScore(list1, list2))
}
