package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

func main() {
	adaptors := Adaptors("input.txt")

	oneDiffs, threeDiffs := countDifferences(adaptors)
	fmt.Printf("Solution 1: %v\n", oneDiffs*threeDiffs)
	fmt.Printf("Solution 2: %v\n", combinations(adaptors, len(adaptors)-1))
}

var cache map[int]int = make(map[int]int)

func combinations(adaptors []int, i int) int {
	if val, ok := cache[i]; ok {
		return val
	}

	if i == 0 {
		return 1
	}

	c := 0
	for j := 1; j <= 3; j++ {
		if i-j >= 0 && adaptors[i-j] >= adaptors[i]-3 {
			c += combinations(adaptors, i-j)
		}
	}
	cache[i] = c
	return c
}

func Adaptors(filename string) []int {
	data, _ := ioutil.ReadFile("input.txt")
	parsed := bytes.Split(data, []byte("\n"))
	adaptors := make([]int, len(parsed))
	for i, line := range parsed {
		adaptors[i], _ = strconv.Atoi(string(line))
	}
	sort.Ints(adaptors)
	return append(append([]int{0}, adaptors...), adaptors[len(parsed)-1]+3)
}

func countDifferences(adaptors []int) (int, int) {
	oneJoltDiffs, threeJoltDiffs := 0, 0
	for i, jolt := range adaptors[1:] {
		diff := jolt - adaptors[i]

		if diff == 1 {
			oneJoltDiffs++
		}
		if diff == 3 {
			threeJoltDiffs++
		}
	}

	return oneJoltDiffs, threeJoltDiffs
}
