package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ExtractIntegersFromFile(filePath string) []int {
	file, err := os.Open(filePath)
	check(err)
	scanner := bufio.NewScanner(file)
	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}
	return depths
}

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v", e)
		os.Exit(1)
	}
}
