package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type timestamp struct {
	year   int
	mont   int
	day    int
	hour   int
	minute int
}

type sleepRegistry struct {
	minuteSleepStart int
	minuteSleepEnd   int
}

func main() {
	bytes, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines) // sort chronologically
	var registroMapa = make(map[string][]sleepRegistry)
	var currentGuard string
	var sleepStart int
	var sleepEnd int
	for _, value := range lines {
		time := parseLine(value)
		message := value[19:]
		messageDestructuring := strings.Fields(message)

		if messageDestructuring[0] == "Guard" {
			currentGuard = messageDestructuring[1][1:]
		}

		if strings.Join(messageDestructuring, " ") == "falls asleep" {
			sleepStart = time.minute
		}

		if strings.Join(messageDestructuring, " ") == "wakes up" {
			sleepEnd = time.minute - 1
		}

		if sleepStart > 0 && sleepEnd > 0 {
			registroMapa[currentGuard] = append(registroMapa[currentGuard], sleepRegistry{
				sleepStart,
				sleepEnd,
			})
			sleepStart = 0
			sleepEnd = 0
		}
	}

	maxMinute, maxGuard := parseMapOfGuards(registroMapa)
	fmt.Println("maxGuard", maxGuard)
	fmt.Println("maxMinute", maxMinute)
	fmt.Println("Resultado final", maxGuard*maxMinute)
}

func parseMapOfGuards(registerMap map[string][]sleepRegistry) (int, int) {
	mapaMinutos := map[string]map[int]int{}
	for key, value := range registerMap {
		if mapaMinutos[key] == nil {
			mapaMinutos[key] = map[int]int{}
		}
		for _, timing := range value {
			for i := timing.minuteSleepStart; i <= timing.minuteSleepEnd; i++ {
				mapaMinutos[key][i]++
			}
		}
	}
	maxMinute, maxGuard := getMaxGuardWithMinute(mapaMinutos)
	return maxMinute, maxGuard
}

func getMaxGuardWithMinute(mapaMinutos map[string]map[int]int) (int, int) {
	mapaFinal := map[int]map[int]int{}
	for id, value := range mapaMinutos {
		if mapaFinal[parseToInt(id)] == nil {
			mapaFinal[parseToInt(id)] = map[int]int{}
		}
		maxMinute, maxTimes := getMaxMinuteAndTimes(value)
		mapaFinal[parseToInt(id)][maxMinute] = maxTimes
	}

	var maxGuard, maxGuardMinute, maxGuardTimes int
	for id, value := range mapaFinal {
		for minute, times := range value {
			if times > maxGuardTimes {
				maxGuard = id
				maxGuardTimes = times
				maxGuardMinute = minute
			}
		}
	}

	return maxGuardMinute, maxGuard
}

func getMaxMinuteAndTimes(mapa map[int]int) (int, int) {
	var maxTimes, maxMinute int
	for minuto, times := range mapa {
		if times > maxTimes {
			maxTimes = times
			maxMinute = minuto
		}
	}

	return maxMinute, maxTimes
}

func parseLine(row string) (parsedLine timestamp) {

	line := strings.Fields(row)
	date := strings.FieldsFunc(line[0], func(r rune) bool {
		switch string(r) {
		case "[", "]", "-":
			return true
		default:
			return false
		}
	})
	time := strings.FieldsFunc(line[1], func(r rune) bool {
		switch string(r) {
		case "]", ":":
			return true
		default:
			return false
		}
	})

	fmt.Println()

	return timestamp{
		parseToInt(date[0]),
		parseToInt(date[1]),
		parseToInt(date[2]),
		parseToInt(time[0]),
		parseToInt(time[1]),
	}
}

func parseToInt(value string) int {
	number, err := strconv.Atoi(strings.Trim(value, " "))

	if err != nil {
		panic(err)
	}

	return number
}
