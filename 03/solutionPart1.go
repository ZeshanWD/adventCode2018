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
	var mapa [1000][1000]int
	for i := 0; i < len(data); i++ {
		claimsArr = append(claimsArr, parseLine(data[i]))
	}

	for _, rect := range claimsArr {
		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.heigth; j++ {
				mapa[i][j]++
			}
		}
	}

	conflicts := countConflicts(mapa)

	fmt.Println("conflicts", conflicts)
}

func countConflicts(mapa [1000][1000]int) int {
	conflicts := 0
	for _, arr := range mapa {
		for _, value := range arr {
			if value > 1 {
				conflicts++
			}
		}
	}

	return conflicts
}

func parseLine(line string) claim {
	item := strings.FieldsFunc(line, func(r rune) bool {
		switch string(r) {
		case "@", ",", "x", ":":
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
