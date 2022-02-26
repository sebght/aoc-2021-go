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

func TestSecondPartTwoWindowsOneIncrease(t *testing.T) {
	result := SecondPart([]int{199, 200, 208, 210})
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestSecondPartTwoWindowsOneDecrease(t *testing.T) {
	result := SecondPart([]int{208, 210, 200, 207})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestSecondPartOfficialExample(t *testing.T) {
	result := SecondPart([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if result != 5 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 5)
	}
}
