package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"bytes"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	commands := bytes.Split(data, []byte("\n"))

	acc, _ := executeProgram(commands)

	fmt.Printf("Solution 1: %v\n", acc)

	i := 0
	for {
		var didFlip bool
		commands, didFlip = flipRow(commands, i)

		if didFlip {
			accResult, didTerminate := executeProgram(commands)
			if didTerminate {
				fmt.Printf("Solution 2: %v\n", accResult)
				break
			}
			commands, _ = flipRow(commands, i)
		}

		i += 1
	}
}

func flipRow(program [][]byte, index int) ([][]byte, bool) {
	line := string(program[index])
	if strings.HasPrefix(line, "jmp") {
		program[index] = append([]byte("nop"), program[index][3:]...)
		return program, true
	} else if strings.HasPrefix(line, "nop") {
		program[index] = append([]byte("jmp"), program[index][3:]...)
		return program, true
	} else {
		return program, false
	}
}

func executeProgram(program [][]byte) (int, bool) {
	i, acc := 0, 0
	history := make(map[int]bool)
	for {
		addI, addAcc := executeCommand(program[i])
		history[i] = true

		i += addI
		acc += addAcc

		if history[i] {
			break
		}

		if i >= len(program) {
			return acc, true
		}
	}

	return acc, false
}

func executeCommand(command []byte) (int, int) {
	cmd := string(command)
	amount, _ := strconv.Atoi(string(command[5:]))

	if strings.HasPrefix(cmd, "acc") {
		if command[4] == '+' {
			return 1, amount
		} else {
			return 1, -amount
		}
    } else if strings.HasPrefix(cmd, "jmp") {
        if command[4] == '+' {
			return amount, 0
		} else {
			return -amount, 0
		}
    } else {
    	return 1, 0
    }
}