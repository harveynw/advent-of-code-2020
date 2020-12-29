package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	count := 0
	position := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		if line[position] == byte(35) { // #
			line[position] = byte(88) // X
			count++
		} else {
			line[position] = byte(79) // O
		}

		fmt.Println(string(line))
		position = (position + 3) % len(line)
	}

	fmt.Printf("Final count: %v\n", count)
}
