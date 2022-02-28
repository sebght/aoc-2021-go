package main

import (
	"aoc-2021-go/days"
	"aoc-2021-go/utils"
	"fmt"
)

func main() {
	fmt.Println("=^..^=    =^..^=   Day 1   =^..^=    =^..^=")
	depths := utils.ExtractIntegersFromFile("inputs/day1.txt")
	fmt.Println("Number of increases :", days.FirstPartDay1(depths))
	fmt.Println("Number of window increases :", days.SecondPartDay1(depths))
	fmt.Print("\n")
	fmt.Println("=^..^=    =^..^=   Day 2   =^..^=    =^..^=")
	lines := utils.ExtractDay2("inputs/day2.txt")
	fmt.Println("Horizontal * Depth :", days.FirstPartDay2(lines))
}
