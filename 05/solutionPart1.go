package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	bytes, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)
	var polimero string
	for scanner.Scan() {
		polimero = scanner.Text()
	}

	cadenaPolimero := getResultPolymer(polimero)
	fmt.Println("length polymer", len(cadenaPolimero))
}

func getResultPolymer(polimero string) string {
	var newString string
	for index, value := range polimero {
		if index+1 < len(polimero) {
			if react(value, rune(polimero[index+1])) {
				fmt.Println("reaccionan", string(value), string(polimero[index+1]))
				newString = polimero[:index] + polimero[index+2:]
				return getResultPolymer(newString)
			}

			newString = polimero
		}
	}

	return newString
}

func react(s1, s2 rune) bool {
	typeCharacter := typeCharacter(string(s1), string(s2))
	polarity := polarity(rune(s1), rune(s2))
	if typeCharacter && polarity {
		return true
	}

	return false
}

func typeCharacter(s1, s2 string) bool {
	if strings.ToUpper(s1) == strings.ToUpper(s2) {
		return true
	}

	return false
}

func polarity(s1, s2 rune) bool {
	if unicode.IsLower(s1) && unicode.IsUpper(s2) || unicode.IsUpper(s1) && unicode.IsLower(s2) {
		return true
	}

	return false
}
