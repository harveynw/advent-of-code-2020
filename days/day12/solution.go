package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"strconv"
)

func main() {
	m := Moves("input.txt")

	b := Boat{'E', 0, 0}
	b.executeMoves(m)

	fmt.Printf("Solution 1: %v\n", abs(b.x) + abs(b.y))

	m = waypointInstructionsToMoves(m)
	b = Boat{'E', 0, 0}
	b.executeMoves(m)

	fmt.Printf("Solution 2: %v\n", abs(b.x) + abs(b.y))
}

type Boat struct {
	direction byte
	x, y int
}

type Move struct {
	direction byte
	n int
}

func waypointInstructionsToMoves(instructions []Move) []Move {
	waypoint := Boat{'N', 10, 1}
	moves := make([]Move, 0)

	for _, v := range instructions {
		if v.direction == 'L' || v.direction == 'R' {
			if v.direction == 'L' {
				v.n = 360 - v.n
			}

			for r := 90; r <= v.n ; r += 90 {
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
			}
		} else if v.direction == 'F' {
			moves = append(moves, Move{'E', v.n * waypoint.x})
			moves = append(moves, Move{'N', v.n * waypoint.y})
		} else {
			waypoint.executeMove(v)
		}
	}

	return moves
}

func (b *Boat) executeMoves(moves []Move) {
	for _, move := range moves {
		b.executeMove(move)
	}
}

func (b *Boat) executeMove(move Move) {
	switch move.direction {
		case 'N':
			b.y += move.n 
		case 'S':
			b.y -= move.n
		case 'E':
			b.x += move.n
		case 'W':
			b.x -= move.n
		case 'L':
			b.turn(360-move.n)
		case 'R':
			b.turn(move.n)
		case 'F':
			b.executeMove(Move{b.direction, move.n})
	}
}

// Turns to the right
func (b *Boat) turn(amount int) {
	next := func() byte {
		switch b.direction {
		case 'N':
			return 'E'
		case 'E':
			return 'S'
		case 'S':
			return 'W'
		case 'W':
			return 'N'
		}
		return 0
	}

	for r := 90; r <= amount; r += 90 {
		b.direction = next()
	}
} 

func Moves(filename string) []Move {
	data, _ := ioutil.ReadFile(filename)
	lines := bytes.Split(data, []byte("\n"))

	moves := []Move{}

	for _, line := range lines {
		n, _ := strconv.Atoi(string(line[1:]))
		moves = append(moves, Move{line[0], n})
	}

	return moves
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}