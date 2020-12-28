package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	expenses, err := getExpenses()
	if err != nil {
		log.Fatalf("Parsing error", err)
	}

	for i, expense_1 := range expenses {
		for j, expense_2 := range expenses {
			if i != j && expense_1 + expense_2 == 2020 {
				fmt.Println(expense_1 * expense_2)
				return
			}
		}
	}
}

func getExpenses() ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var expenses []int
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}