package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bytes, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)

	var numberArr []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numberArr = append(numberArr, i)
	}

	var newArr []int
	resultValue := 0
	for true {
		for _, currentValue := range numberArr {
			resultValue += currentValue
			if Contains(newArr, resultValue) {
				fmt.Println("found: ", resultValue)
				os.Exit(0)
			}
			newArr = append(newArr, resultValue)
		}
	}
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
