package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"bytes"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validCount := 0

	for _, passData := range bytes.Split(data, []byte("\n\n")) {
		passData = bytes.ReplaceAll(passData, []byte("\n"), []byte(" "))
		fields := bytes.Split(passData, []byte(" "))

		passport := make(map[string]string) 
		for _, field := range fields {
			parsed := bytes.Split(field, []byte(":"))
			if len(parsed) == 2 {
				passport[string(parsed[0])] = string(parsed[1])
			}
		}

		fmt.Println(passport)

		if isValid(passport) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func isValid(passport map[string]string) bool {
	required := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, requiredField := range required {
		if _, ok := passport[requiredField]; !ok {
			return false
		}
	}

	return true
}