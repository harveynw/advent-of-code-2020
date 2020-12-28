package main 

import (
	"fmt"
	"log"
	"bytes"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(data, []byte("\n\n"))

	count := 0
	for _, line := range lines {
		choices := make(map[byte]bool)
		for _, b := range bytes.ReplaceAll(line, []byte("\n"), []byte("")) {
			choices[b] = true
		}
		count += len(choices)
	}

	fmt.Printf("Solution 1: %v\n", count)

	count = 0
	for _, line := range lines {
		choices := make(map[byte]int)

		people := bytes.Split(line, []byte("\n"))
		threshold := len(people)

		for i, person := range people {
			for _, b := range person {
				choices[b]++

				if i == threshold - 1 && choices[b] == threshold {
					count++
				}
			}
		}
	}

	fmt.Printf("Solution 2: %v\n", count)
}