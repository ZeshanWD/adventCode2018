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

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		fmt.Println("Number: ", i)
		count += i
	}

	fmt.Println("la suma final es: ", count)
}
