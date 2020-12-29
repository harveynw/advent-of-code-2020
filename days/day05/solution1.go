package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	seats := bytes.Split(data, []byte("\n"))

	highestSeat := seats[0]
	for _, seat := range seats {
		if compareSeats(seat, highestSeat) {
			fmt.Printf("%v > %v\n", string(seat), string(highestSeat))
			highestSeat = seat
		}
	}

	fmt.Println(string(highestSeat))
	fmt.Println(computeSeatID(highestSeat))
}

func compareSeats(seatA []byte, seatB []byte) bool {
	for i, v := range seatA {
		if i < 7 {
			if v < seatB[i] {
				return true
			}
			if v > seatB[i] {
				return false
			}
		} else {
			if v > seatB[i] {
				return true
			}
			if v < seatB[i] {
				return false
			}
		}
	}
	return false
}

func computeSeatID(seat []byte) int {
	row := 0

	for i, val := range seat[0:7] {
		if val == 'B' {
			row += int(math.Pow(2.0, float64(6-i)))
		}
	}

	column := 0

	for i, val := range seat[7:10] {
		if val == 'R' {
			column += int(math.Pow(2.0, float64(2-i)))
		}
	}

	return row*8 + column
}
