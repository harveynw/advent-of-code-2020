package main

import (
	"fmt"
	"log"
	"bytes"
	"io/ioutil"
	"math/rand"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	seats := parseSeats(data)
	seats = quicksort(seats)

	for i, _ := range seats {
		if computeSeatID(seats[i+1])-computeSeatID(seats[i])>1 {
			fmt.Println(computeSeatID(seats[i])+1)
			return
		}
	}
}

func parseSeats(data []byte) [][]bool {
	lines := bytes.Split(data, []byte("\n"))

	seats := make([][]bool, len(lines))
	for i := 0; i < len(lines); i++ {
	    seats[i] = make([]bool, 10)
	}

	for i, line := range lines {
		for j, char := range line {
			seats[i][j] = (char == 'B' || char == 'R')
		}
	}

	return seats
}

func quicksort(seats [][]bool) [][]bool {
	if len(seats) < 2 {
		return seats
	}

	compare := func(s1 []bool, s2 []bool) bool {
		for i := 0; i < 10; i++ {
			if s1[i] != s2[i] {
				return s2[i]
			}
		} 
		return false
	}

	l, r := 0, len(seats) - 1
	pivot := rand.Int() % len(seats)

	seats[pivot], seats[r] = seats[r], seats[pivot]

	for i, _ := range seats {
		if compare(seats[i], seats[r]) {
			seats[l], seats[i] = seats[i], seats[l]
			l++
		}
	}

	seats[l], seats[r] = seats[r], seats[l]

	quicksort(seats[:l])
	quicksort(seats[l+1:])

	return seats
}

func computeSeatID(seat []bool) int {
	id := 0
	pow := 512
	for _, val := range seat[0:10] {
		if val {
			id += pow
		}
		pow /= 2
	}
	return id
}
