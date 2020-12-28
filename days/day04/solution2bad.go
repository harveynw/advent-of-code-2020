package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"bytes"
	"strconv"
	"strings"
	"regexp"
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
			fmt.Printf("Missing %v\n", requiredField)
			return false
		}
	}

	byr, _ := strconv.Atoi(passport["byr"])
	iyr, _ := strconv.Atoi(passport["iyr"])

	if byr < 1920 || byr > 2002 || iyr < 2010 || iyr > 2020 {
		fmt.Println("byr")
		return false
	}

	eyr, _ := strconv.Atoi(passport["eyr"])

	if eyr < 2020 || eyr > 2030 {
		fmt.Println("eyr")
		return false
	}

	hgt := passport["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		val, _ := strconv.Atoi(strings.ReplaceAll(hgt, "cm", ""))
		if val < 150 || val > 193 {
			fmt.Println("hgt_cm")
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		val, _ := strconv.Atoi(strings.ReplaceAll(hgt, "in", ""))
		if val < 59 || val > 76 {
			fmt.Println("hgt_in")
			return false
		}
	} else {
		fmt.Println("hgt")
		return false
	}

	hclFormat, _ := regexp.Compile("^#[0-9a-f]{6}$")
	if !hclFormat.Match([]byte(passport["hcl"])) {
		fmt.Println("hcl")
		return false
	}

	ecl := passport["ecl"]
	validColours := map[string]bool {
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if !validColours[ecl] {
		fmt.Println("ecl")
		return false
	}

	pidFormat, _ := regexp.Compile("^[0-9]{9}$")
	if !pidFormat.Match([]byte(passport["pid"])) {
		fmt.Println("pid")
		return false
	}

	return true
}