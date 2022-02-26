package main

import (
	"aoc-2021-go/days"
	"aoc-2021-go/utils"
	"fmt"
)

func main() {
	fmt.Println("=^..^=    =^..^=  Day 1  =^..^=  =^..^=")
	depths := utils.ExtractIntegersFromFile("inputs/day1.txt")
	fmt.Println("Number of increases :", days.FirstPart(depths))
	fmt.Println("Number of window increases :", days.SecondPart(depths))
}
