package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"strconv"
)

func main() {
	buses, diffs := Buses("input.txt")
	fmt.Println(buses, diffs)
	fmt.Println(solveCongruences(diffs, buses))
}

/*
	Solves the system of congruences:
	t ≡ a_i mod m_i
	Under the assumption that {m_1, ..., m_n} are pairwise coprime
*/
func solveCongruences(a, m []int) int {
	M := product(m)

	M_i, y := make([]int, len(m)), make([]int, len(m))

	for i, _ := range a {
		M_i[i] = M / m[i]
		gcd, y_i, _ := egcd(M_i[i], m[i])
		y[i] = lpr(y_i, gcd)
	}

	r := sumProduct(a, y, M_i)

	// for r < 0 {
	// 	r += M
	// }

	fmt.Println(M)
	return r
}

func qr(a, b int) (quotient, remainder int) {
	quotient, remainder = a / b, a % b
	return
}

func egcd(a, b int) (gcd, x, y int) {
	x, y = 0, 1
	u, v := 1, 0
	for a != 0 {
		q, r := qr(b, a)
		m, n := x - u*q, y - v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	gcd = b
	return
}

func lpr(a, m int) int {
	return (a % m + m) % m
}

func product(list []int) int {
	p := 1
	for _, v := range list {
		p *= v
	}
	return p
}

func sumProduct(lists ...[]int) int {
	result := 0
	for i, _ := range lists[0] {
		p := 1
		for _, list := range lists {
			p *= list[i]
		}
		result += p
	}
	return result
}


func Buses(filename string) (buses, diffs []int) {
	data, _ := ioutil.ReadFile(filename)
	lines := bytes.Split(data, []byte("\n"))

	for i, n := range bytes.Split(lines[1], []byte(",")) {
		if id := string(n); id != "x" {
			intID, _ := strconv.Atoi(id)
			buses, diffs = append(buses, intID), append(diffs, i)
		}
	}

	return
}
