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
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	var claimsArr []claim
	var cleanArr = make(map[string]int)
	var mapa [1000][1000]int
	for i := 0; i < len(data); i++ {
		lineaParseada := parseLine(data[i])
		claimsArr = append(claimsArr, lineaParseada)
	}

	for _, rect := range claimsArr {
		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.heigth; j++ {
				mapa[i][j]++
				cleanArr[rect.id] = mapa[i][j]
			}
		}
	}

	// conflicts := countConflicts(mapa)

	fmt.Println("clean claims", getCleanClaim(cleanArr))
}

func getCleanClaim(arr map[string]int) string {
	var found string
	for key, _ := range arr {
		if arr[key] == 1 {
			found = key
			fmt.Println("Found", found)
			os.Exit(0)
		}
	}

	return found
}

// func countConflicts(mapa [1000][1000]int) int {
// 	conflicts := 0
// 	for _, arr := range mapa {
// 		for _, value := range arr {
// 			if value > 1 {
// 				conflicts++
// 			}
// 		}
// 	}

// 	return conflicts
// }

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
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
