package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)

	twoLetters := 0
	threeLetter := 0
	for scanner.Scan() {
		line := scanner.Text()
		containsWords(line, &twoLetters, &threeLetter)
	}
	fmt.Println("Final result: ", twoLetters*threeLetter)
}

func containsWords(line string, twoLetters *int, threeLetter *int) {
	var count int
	var countsArr []int
	for _, char := range line {
		count = strings.Count(line, string(char))
		if count == 2 {
			countsArr = append(countsArr, count)
		}
		if count == 3 {
			countsArr = append(countsArr, count)
		}
	}

	if hasNumber(countsArr, 2) {
		*twoLetters++
	}

	if hasNumber(countsArr, 3) {
		*threeLetter++
	}
}

func hasNumber(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
