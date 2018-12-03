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
	var wordsArr []string
	for scanner.Scan() {
		wordsArr = append(wordsArr, scanner.Text())
	}

	for _, value := range wordsArr {
		analyseId(value, wordsArr)
	}
}

func analyseId(value string, wordsArr []string) {
	for _, currentValue := range wordsArr {
		compareString(value, currentValue)
	}
}

func compareString(string_1, string_2 string) {
	var resultArr []bool
	var resultChars []string
	for i := 0; i < len(string_1); i++ {
		if string(string_1[i]) == string(string_2[i]) {
			resultArr = append(resultArr, true)
			resultChars = append(resultChars, string(string_1[i]))
		}
		resultArr = append(resultArr, false)
	}

	if len(resultChars) == (len(string_1) - 1) {
		fmt.Println(strings.Join(resultChars[:], ""))
		os.Exit(0)
	}
}
