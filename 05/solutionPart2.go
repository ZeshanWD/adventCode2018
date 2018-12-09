package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	bytes, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)
	var polimero string
	for scanner.Scan() {
		polimero = scanner.Text()
	}

	minLengPolymer := getMinimumLengthPolymer(polimero)
	fmt.Println("length polymer", minLengPolymer)
}

func getMinimumLengthPolymer(cadenaPolimero string) int {
	var stringWithRemovedChar string
	var fullyReactedString string
	lettersMap := map[string]int{}
	for _, value := range cadenaPolimero {
		fmt.Println("letter", string(value))
		stringWithRemovedChar = removeCharactersFromString(string(value), cadenaPolimero)
		// fmt.Println("removed Letter string", stringWithRemovedChar)
		fullyReactedString = getResultPolymer(stringWithRemovedChar)
		lettersMap[string(value)] = len(fullyReactedString)
	}

	fmt.Println("Mapa", lettersMap)

	return getMaxLength(lettersMap)
}

func removeCharactersFromString(character string, cadena string) string {
	var newString string
	fmt.Println("previous", cadena)
	newString = strings.Replace(cadena, strings.ToUpper(string(character)), "", -1)
	newString = strings.Replace(cadena, strings.ToLower(string(character)), "", -1)
	fmt.Println("after", newString)
	return newString
}

func getMaxLength(mapa map[string]int) int {
	var maxKey string
	var maxLength int
	for key, value := range mapa {
		if value < maxLength {
			maxLength = value
			maxKey = key
		}
	}
	fmt.Println("maxKey", maxKey)
	fmt.Println("maxLentgh", maxLength)
	return maxLength
}

func getResultPolymer(polimero string) string {
	var newString string
	for index, value := range polimero {
		if index+1 < len(polimero) {
			if react(value, rune(polimero[index+1])) {
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
