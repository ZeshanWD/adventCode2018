package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("Final result: ", wordsArr)
}
