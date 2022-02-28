package days

import (
	"strconv"
	"strings"
)

func FirstPartDay2(lines []string) int {
	horizontal := 0
	depth := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])
		switch direction {
		case "forward":
			depth += distance
		case "up":
			horizontal -= distance
		case "down":
			horizontal += distance
		}
	}
	return horizontal * depth
}

func SecondPartDay2(lines []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])
		switch direction {
		case "forward":
			horizontal += distance
			depth += distance * aim
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return horizontal * depth
}
