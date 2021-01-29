package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"strconv"
)

func main() {
	earliestTime, buses := Buses("input.txt")
	
	i, wait := earliestBus(buses, earliestTime)

	fmt.Printf("Solution 1: %v\n", buses[i] * wait)

}

func earliestBus(buses []int, earliestTime int) (int, int) {
	i, minWait := 0, buses[0] - earliestTime % buses[0]

	for j, v := range buses[1:] {
		if wait := v - earliestTime % v; wait < minWait {
			i, minWait = j+1, wait
		}
	}

	return i, minWait
}


func Buses(filename string) (int, []int) {
	data, _ := ioutil.ReadFile(filename)
	lines := bytes.Split(data, []byte("\n"))

	earliest, _ := strconv.Atoi(string(lines[0]))

	buses := make([]int, 0)
	for _, n := range bytes.Split(lines[1], []byte(",")) {
		if id := string(n); id != "x" {
			intID, _ := strconv.Atoi(id)
			buses = append(buses, intID)
		}
	}

	return earliest, buses
}
