package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadDay1(filePath string) []int {
	file, err := os.Open(filePath)
	Check(err)
	scanner := bufio.NewScanner(file)
	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}
	return depths
}

func ReadDay2(filePath string) []string {
	file, err := os.Open(filePath)
	Check(err)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
