package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type Bag struct {
	name     string
	children []BagRequirement
}

type BagRequirement struct {
	name   string
	amount int
}

func (b *Bag) containsBag(name string) bool {
	for _, child := range b.children {
		if name == child.name {
			return true
		}
	}
	return false
}

func main() {
	bags := FetchBags()

	_, count := countParents(bags, "shiny gold")

	fmt.Printf("Solution 1: %v\n", count)

	bags = FetchBags()

	count = countDescendents(bags, "shiny gold")
	fmt.Printf("Solution 2: %v\n", count-1)
}

func FetchBags() []Bag {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var bags []Bag

	for _, line := range bytes.Split(data, []byte("\n")) {
		bags = append(bags, ParseBag(line))
	}

	return bags
}

func countDescendents(bags []Bag, target string) int {
	count := 0
	var newTargets []BagRequirement

	for i := 0; i < len(bags); i++ {
		if bags[i].name == target {
			newTargets = bags[i].children
			break
		}
	}

	var descendentCount int
	for _, newTarget := range newTargets {
		descendentCount = countDescendents(bags, newTarget.name)
		count += newTarget.amount * descendentCount
	}

	return count + 1
}

func countParents(bags []Bag, target string) ([]Bag, int) {
	count := 0

	var newTargets []string

	removeElement := func(i int) {
		bags[i] = bags[len(bags)-1]
		bags = bags[:len(bags)-1]
	}

	for i := len(bags) - 1; i >= 0; i-- {
		if bags[i].containsBag(target) {
			newTargets = append(newTargets, bags[i].name)
			removeElement(i)
			count++
		}
	}

	var increment int
	if len(bags) > 0 {
		for _, newTarget := range newTargets {
			bags, increment = countParents(bags, newTarget)
			count += increment
		}
	}

	return bags, count
}

var bagNamePattern, _ = regexp.Compile(`^(\d{1,2}) ([\w\s]+) bag`)

func ParseBag(line []byte) Bag {
	parts := bytes.Split(line, []byte(" bags contain "))

	name := string(parts[0])
	var children []BagRequirement

	children_defns := bytes.Split(parts[1], []byte(", "))
	if string(children_defns[0]) != "no other bags." {
		for _, child_defn := range children_defns {
			vals := bagNamePattern.FindSubmatch(child_defn)

			number, _ := strconv.Atoi(string(vals[1]))
			colour := string(vals[2])

			r := BagRequirement{colour, number}
			children = append(children, r)
		}
	}

	return Bag{name, children}
}
