package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	multiple := 1
	multiple *= testSlope([2]int{1, 1}, file)
	multiple *= testSlope([2]int{3, 1}, file)
	multiple *= testSlope([2]int{5, 1}, file)
	multiple *= testSlope([2]int{7, 1}, file)
	multiple *= testSlope([2]int{1, 2}, file)

	fmt.Printf("Final value: %v\n", multiple)
}

func testSlope(slope [2]int, file *os.File) int {
	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)

	var count int = 0
	var position int = 0

	for scanner.Scan() {
		line := scanner.Bytes()

		if line[position] == byte(35) { // #
			count++
		}

		position = (position + slope[0]) % len(line)

		for i := 1; i < slope[1]; i++ {
			scanner.Scan()
		}
	}

	return count
}
