package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"strconv"
	"math/rand"
	"math"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	parsed := bytes.Split(data, []byte("\n"))
	adaptors := make([]int, len(parsed))
	for i, line := range parsed {
		adaptors[i], _ = strconv.Atoi(string(line))
	}
	adaptors = quicksort(adaptors)
	
	oneDiffs, threeDiffs := countDifferences(adaptors)
	fmt.Printf("Solution 1: %v\n", oneDiffs*threeDiffs)
	fmt.Printf("Solution 2: %v\n", countCombinations(adaptors))
}

// Divide and conquer (just the once)
func countCombinations(adaptors []int) int {
	droppable := droppableAdaptors(adaptors)

	combinations := 1

	for i := 0; i < len(droppable) - 1; i++ {
		if droppable[i] {
			start := int(i)
			for droppable[i] {
				i++
			}
			end := int(i)
			// Break it down into sequential droppable subsets
			// Then naively test and count combinations
			combinations *= countCombinationsSubset(adaptors, start, end)
		}
	}

	return combinations
}

func droppableAdaptors(adaptors []int) []bool {
	droppable := make([]bool, len(adaptors))
	for i := 0; i < len(adaptors) - 1; i++ {
		if i == 0 {
			droppable[i] = adaptors[1] <= 3
		} else {
			droppable[i] = adaptors[i-1] + 3 >= adaptors[i+1]
		}
	}
	return droppable
}

func countCombinationsSubset(adaptors []int, start int, end int) int {
	multiplier := 0

	n := end - start

	test := make([]bool, n)
	for j := 0; j < powInt(2, n); j++ {
		itob(test, j)
		var currentJoltage int
		valid := true
		if start == 0 {
			currentJoltage = 0
		} else {
			currentJoltage = adaptors[start-1]
		}
		for i, drop := range test {
			if drop {
				if currentJoltage + 3 < adaptors[start+i+1] {
					valid = false
					break
				}
			} else {
				currentJoltage = adaptors[start+i]
			}
		}

		if valid {
			multiplier++
		}
	}

	return multiplier

}

func itob(b []bool, n int) {
	for i, _ := range b {
		b[i] = !(n|(1<<i) > n)
	}
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	l, r := 0, len(arr) - 1
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[r] = arr[r], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[r] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]

	quicksort(arr[:l])
	quicksort(arr[l+1:])

	return arr
}

func countDifferences(adaptors []int) (int, int) {
	oneJoltDiffs, threeJoltDiffs := 1,1
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
