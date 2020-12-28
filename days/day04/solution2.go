package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
	"bytes"
	"regexp"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	v := CreateValidator()
	v.addConstraint("byr", func(value string) bool {
		i, _ := strconv.Atoi(value)
		return i >= 1920 && i <= 2002
	})
	v.addConstraint("iyr", func(value string) bool {
		i, _ := strconv.Atoi(value)
		return i >= 2010 && i <= 2020
	})
	v.addConstraint("eyr", func(value string) bool {
		i, _ := strconv.Atoi(value)
		return i >= 2020 && i <= 2030
	})
	v.addConstraint("hgt", func(value string) bool {
		if strings.HasSuffix(value, "cm") {
			i, _ := strconv.Atoi(strings.ReplaceAll(value, "cm", ""))
			return i >= 150 && i <= 193
		} else if strings.HasSuffix(value, "in") {
			i, _ := strconv.Atoi(strings.ReplaceAll(value, "in", ""))
			return i >= 59 && i <= 76
		} else {
			return false
		}
	})
	hclFormat, _ := regexp.Compile("^#[0-9a-f]{6}$")
	v.addConstraint("hcl", func(value string) bool {
		return hclFormat.Match([]byte(value))
	})
	v.addConstraint("ecl", func(value string) bool {
		valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, colour := range valid {
			if colour == value {
				return true
			}
		}
		return false
	})
	pidFormat, _ := regexp.Compile("^[0-9]{9}$")
	v.addConstraint("pid", func(value string) bool {
		return pidFormat.Match([]byte(value))
	})

	validCount := 0
	for _, passData := range bytes.Split(data, []byte("\n\n")) {
		passport := ParsePassport(passData)
		if v.validate(passport) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type Passport struct {
	fields map[string]string
}

type FieldValidator struct {
	constraints map[string]func(string) bool
}

func (validator *FieldValidator) addConstraint(field string, constraint func(string) bool) {
	validator.constraints[field] = constraint
}

func (validator *FieldValidator) validate(p Passport) bool {
	required := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, requiredField := range required {
		if _, ok := p.fields[requiredField]; !ok {
			return false
		}
	}

	for key, constraint := range validator.constraints {
		if !constraint(p.fields[key]) {
			return false
		}
	}

	return true
}

func ParsePassport(definition []byte) Passport {
	data := bytes.ReplaceAll(definition, []byte("\n"), []byte(" "))
	fieldsData := bytes.Split(data, []byte(" "))

	fields := make(map[string]string)
	for _, field := range fieldsData {
		parsed := bytes.Split(field, []byte(":"))
		fields[string(parsed[0])] = string(parsed[1])
	}

	return Passport{fields}
}

func CreateValidator() FieldValidator {
	return FieldValidator{make(map[string]func(string) bool)}
}
