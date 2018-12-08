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
	id := getMostSleepGuard(registroMapa)
	minuteMostSleeped := getMinuteMostSleeped(registroMapa, id)
	fmt.Println("Guard Id", id)
	fmt.Println("Minute most sleeped", minuteMostSleeped)

	fmt.Println(parseToInt(id) * minuteMostSleeped)
}

func getMinuteMostSleeped(registroMapa map[string][]sleepRegistry, guardId string) int {
	guard := registroMapa[string(guardId)]
	resultMap := map[int]int{}
	for _, value := range guard {
		resultMap[value.minuteSleepStart]++
		resultMap[value.minuteSleepEnd]++
	}

	return getMaxNumberRepetitions(resultMap)
}

func getMaxNumberRepetitions(mapa map[int]int) int {

	var maxKey int
	var maxValue int
	for key, value := range mapa {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	return int(maxKey)
}

func getMostSleepGuard(registroMapa map[string][]sleepRegistry) string {
	mapaMostMinutesSleep := map[string]int{}
	var totalMinutes int
	for key, value := range registroMapa {
		for _, minuteRegistry := range value {
			totalMinutes += (minuteRegistry.minuteSleepEnd + 1) - minuteRegistry.minuteSleepStart
		}
		mapaMostMinutesSleep[string(key)] = totalMinutes
		totalMinutes = 0
	}

	return getMaxMinutesGuard(mapaMostMinutesSleep)
}

func getMaxMinutesGuard(mapa map[string]int) string {

	var maxKey string
	var maxMinutes int
	for key, value := range mapa {
		if value > maxMinutes {
			maxMinutes = value
			maxKey = key
		}
	}

	return maxKey
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
