package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
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
	for scanner.Scan() {
		if isLineValid(scanner.Bytes()) {
			count++
		}
	}

	fmt.Println(count)
}

func isLineValid(line []byte) bool {
	split := bytes.Split(line, []byte(" "))
	bounds := bytes.Split(split[0], []byte("-"))

	min, _ := strconv.Atoi(string(bounds[0]))
	max, _ := strconv.Atoi(string(bounds[1]))
	character := split[1][0]
	password := split[2]

	p := Policy{character, min, max}

	return p.isCompliant(password)
}

type Policy struct {
	character byte
	min       int
	max       int
}

func (p Policy) isCompliant(password []byte) bool {
	count := 0
	for _, val := range password {
		if val == p.character {
			count++
		}
	}

	return count >= p.min && count <= p.max
}
