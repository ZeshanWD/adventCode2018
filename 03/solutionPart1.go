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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
