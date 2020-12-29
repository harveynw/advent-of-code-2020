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

	idx_1, _ := strconv.Atoi(string(bounds[0]))
	idx_2, _ := strconv.Atoi(string(bounds[1]))
	character := split[1][0]
	password := split[2]

	p := Policy{character, idx_1, idx_2}

	return p.isCompliant(password)
}

type Policy struct {
	character byte
	idx_1     int
	idx_2     int
}

func (p Policy) isCompliant(password []byte) bool {
	var test1 bool = password[p.idx_1-1] == p.character
	var test2 bool = password[p.idx_2-1] == p.character

	return test1 != test2
}
