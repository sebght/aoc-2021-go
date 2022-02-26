package days

import (
	"aoc-2021-go/utils"
	"testing"
)

func TestFirstPartOneIncrease(t *testing.T) {
	result := FirstPart([]int{1, 2})
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestFirstPartOneDecrease(t *testing.T) {
	result := FirstPart([]int{2, 1})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestFirstPartMultipleIncreases(t *testing.T) {
	result := FirstPart([]int{1, 2, 4, 3})
	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 2)
	}
}

func TestFirstPartRealData(t *testing.T) {
	result := FirstPart(utils.ExtractIntegersFromFile("../inputs/day1.txt"))
	if result != 1228 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1228)
	}
}
