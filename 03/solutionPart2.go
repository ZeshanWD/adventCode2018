package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     string
	x      int
	y      int
	width  int
	heigth int
}

func main() {
	bytes, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)
	var claimsArr []claim
	var mapa [1000][1000]int
	for scanner.Scan() {
		claimsArr = append(claimsArr, parseLine(scanner.Text()))
	}

	for _, rect := range claimsArr {
		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.heigth; j++ {
				mapa[i][j]++
			}
		}
	}

	for _, claim := range claimsArr {
		if isCleanClaim(mapa, claim) {
			fmt.Println("Clean Claim found: ", claim.id)
		}
	}
}

func isCleanClaim(arr [1000][1000]int, claim claim) bool {
	for i := claim.x; i < claim.x+claim.width; i++ {
		for j := claim.y; j < claim.y+claim.heigth; j++ {
			if arr[i][j] > 1 {
				return false
			}
		}
	}

	return true
}

func parseLine(line string) claim {
	item := strings.FieldsFunc(line, func(r rune) bool {
		switch string(r) {
		case "@", ",", "x", ":", "#":
			return true
		default:
			return false
		}
	})

	return claim{
		item[0],
		parseToInt(item[1]),
		parseToInt(item[2]),
		parseToInt(item[3]),
		parseToInt(item[4]),
	}
}

func parseToInt(value string) int {
	number, err := strconv.Atoi(strings.Trim(value, " "))

	if err != nil {
		panic(err)
	}

	return number
}
