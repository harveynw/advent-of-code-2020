package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
)

func main() {
	seats := Seats("input.txt")

	for { 
		count, next := iterate(seats)
		seats = next

		if count == 0 {
			fmt.Printf("Solution 2: %v\n", countOccupied(seats))
			break
		}
	}
}

func countOccupied(seats [][]byte) int {
	count := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[0]); j++ {
			if seats[i][j] == Taken {
				count++
			}
		}
	}
	return count
}

func iterate(seats [][]byte) (int, [][]byte) {
	dimI, dimJ := len(seats), len(seats[0])
	changed := 0
	newSeats := allocateNewSeats(dimI, dimJ)

	for i := 0; i < dimI; i++ {
		for j := 0; j < dimJ; j++ {
			didChange, state := flipSeat(i, j, dimI, dimJ, seats)

			if didChange {
				changed++
			}

			newSeats[i][j] = state
		}
	}

	return changed, newSeats
}

func countTaken(i, j, dimI, dimJ int, seats[][]byte) int {
	taken := 0

	directions := []int{1,0,1,1,0,1,-1,0,-1,-1,0,-1,1,-1,-1,1}

	for dir := 0; dir < len(directions); dir += 2 {
		curI, curJ := i, j

		for {
			curI += directions[dir]
			curJ += directions[dir+1]

			if curI < 0 || curJ < 0 || curI == dimI || curJ == dimJ {
				break
			}

			if seats[curI][curJ] == Taken {
				taken++
				break
			} else if seats[curI][curJ] == Free {
				break
			}
		}
	}

	return taken
}

func flipSeat(i, j, dimI, dimJ int, seats [][]byte) (bool, byte) {
	switch seats[i][j] {
	case Free:
		occupied := countTaken(i, j, dimI, dimJ, seats)

		if occupied > 0 {
			return false, Free
		} else {
			return true, Taken
		}

	case Taken:
		occupied := countTaken(i, j, dimI, dimJ, seats)

		if occupied >= 5 {
			return true, Free
		} else {
			return false, Taken
		}
	}
	return false, Floor
}

const Taken, Free, Floor byte = '#', 'L', '.'

func Seats(filename string) [][]byte {
	data, _ := ioutil.ReadFile(filename)
	lines := bytes.Split(data, []byte("\n"))

	seats := allocateNewSeats(len(lines), len(lines[0]))
	for i, line := range lines {
		seats[i] = line
	}

	return seats
}

func allocateNewSeats(dimI, dimJ int) [][]byte {
	seats := make([][]byte, dimI)
	for i, _ := range seats {
		seats[i] = make([]byte, dimJ)
	}
	return seats
}

func min(a, b int) int {
	if a < b {
		return a 
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}