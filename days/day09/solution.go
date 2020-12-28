package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"bytes"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := bytes.Split(data, []byte("\n"))

	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(string(line))
	}

	var solution int
	for i := 25; i < len(numbers); i++ {
		if !valid(i, numbers) {
			solution = numbers[i]
			break
		}
	}

	setLength := 2
	for {
		for i := 0; i + setLength - 1 < len(numbers); i++ {
			if contiguousSum(i, i + setLength - 1, numbers) == solution {
				min, max := minMaxIntSlice(numbers[i:i+setLength])
				fmt.Printf("Solution 1: %v\n", solution)
				fmt.Printf("Solution 2: %v\n", min + max)
				return
			}
		}
		setLength++
	}
}

func valid(idx int, numbers []int) bool {
	preamble := numbers[max(idx-25, 0):idx]
	l := len(preamble)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if preamble[i] + preamble[j] == numbers[idx] {
				return true
			}
		}
	}

	return false
}

func minMaxIntSlice(input []int) (int, int) {
	min, max := input[0], input[0]

	for _, v := range input {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return min, max
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func contiguousSum(i int, j int, numbers []int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += numbers[k]
	}
	return sum
}
