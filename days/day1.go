package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v", e)
		os.Exit(1)
	}
}

func First_part(file_path string) int {
	depths := extract_integers_from_file(file_path)
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			count++
		}
	}
	return count
}

func extract_integers_from_file(file_path string) []int {
	file, err := os.Open(file_path)
	check(err)
	scanner := bufio.NewScanner(file)
	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}
	return depths
}
